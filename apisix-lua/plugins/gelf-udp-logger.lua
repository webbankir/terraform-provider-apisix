local core = require("apisix.core")
local log_util = require("apisix.utils.log-util")
local bp_manager_mod = require("apisix.utils.batch-processor-manager")
local plugin_name = "udp-logger"
local tostring = tostring
local ngx = ngx
local udp = ngx.socket.udp


local batch_processor_manager = bp_manager_mod.new("udp logger")
local schema = {
    type = "object",
    properties = {
        host = {type = "string"},
        port = {type = "integer", minimum = 0},
        timeout = {type = "integer", minimum = 1, default = 3},
        include_req_body = {type = "boolean", default = false}
    },
    required = {"host", "port"}
}


local _M = {
    version = 0.1,
    priority = 401,
    name = plugin_name,
    schema = batch_processor_manager:wrap_schema(schema),
}

function _M.check_schema(conf)
    return core.schema.check(schema, conf)
end

local function get_full_log(ngx, conf)
    local ctx = ngx.ctx.api_ctx
    local var = ctx.var
    local service_id
    local route_id
    local url = var.scheme .. "://" .. var.host .. ":" .. var.server_port
            .. var.request_uri
    local matched_route = ctx.matched_route and ctx.matched_route.value

    if matched_route then
        service_id = matched_route.service_id or ""
        route_id = matched_route.id
    else
        service_id = var.host
    end

    local consumer
    if ctx.consumer then
        consumer = {
            username = ctx.consumer.username
        }
    end

    local log = {
        request = {
            url = url,
            uri = var.request_uri,
            method = ngx.req.get_method(),
            headers = ngx.req.get_headers(),
            querystring = ngx.req.get_uri_args(),
            size = var.request_length
        },
        response = {
            status = ngx.status,
            headers = ngx.resp.get_headers(),
            size = var.bytes_sent
        },
        server = {
            hostname = core.utils.gethostname(),
            version = core.version.VERSION
        },
        upstream = var.upstream_addr,
        service_id = service_id,
        route_id = route_id,
        consumer = consumer,
        client_ip = core.request.get_remote_client_ip(ngx.ctx.api_ctx),
        start_time = ngx.req.start_time() * 1000,
        latency = (ngx_now() - ngx.req.start_time()) * 1000,
        log_time_iso = var.time_iso8601
    }

    if ctx.resp_body then
        log.response.body = ctx.resp_body
    end

    if conf.include_req_body then

        local log_request_body = true

        if conf.include_req_body_expr then

            if not conf.request_expr then
                local request_expr, err = expr.new(conf.include_req_body_expr)
                if not request_expr then
                    core.log.error('generate request expr err ' .. err)
                    return log
                end
                conf.request_expr = request_expr
            end

            local result = conf.request_expr:eval(ctx.var)

            if not result then
                log_request_body = false
            end
        end

        if log_request_body then
            local body = req_get_body_data()
            if body then
                log.request.body = body
            else
                local body_file = ngx.req.get_body_file()
                if body_file then
                    log.request.body_file = body_file
                end
            end
        end
    end

    return {
        version = "1.1",
        host = ngx.var.hostname,
        short_message = "[" .. ngx.req.get_method() .. "] " .. url .. " [" .. ngx.status .. "]",
        full_message = json.encode(log)
    }
end

function _M.log(conf, ctx)
    local entry = get_full_log(ngx, conf)

    if batch_processor_manager:add_entry(conf, entry) then
        return
    end

    -- Generate a function to be executed by the batch processor
    local func = function(entries, batch_max_size)
        local err_msg
        local res = true
        local sock = udp()
        sock:settimeout(conf.timeout * 1000)
        local ok, err = sock:setpeername("graylog.team.webbankir.cloud", 21002)
        if not ok then
            return false, "failed to connect to UDP server: host[" .. conf.host
                    .. "] port[" .. tostring(conf.port) .. "] err: " .. err
        end

        for _, entry in ipairs(entries) do
            local data, err = core.json.encode(entry)
            if not data then
                return false, 'error occurred while encoding the data: ' .. err
            end

            ok, err = sock:send(data)
            if not ok then
                res = false
                err_msg = "failed to send data to UDP server: host[" .. conf.host
                        .. "] port[" .. tostring(conf.port) .. "] err:" .. err
            end
            ok, err = sock:close()
            if not ok then
                core.log.error("failed to close the UDP connection, host[",
                        conf.host, "] port[", conf.port, "] ", err)
            end

            return res, err_msg
        end
    end

    batch_processor_manager:add_entry_to_new_processor(conf, entry, ctx, func)
end

return _M