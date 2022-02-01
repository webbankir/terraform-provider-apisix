local batch_processor = require("apisix.utils.batch-processor")
local core = require("apisix.core")
local http = require("resty.http")
local json = require('cjson')
local ngx = ngx
local ngx_now = ngx.now
local tostring = tostring
local ipairs = ipairs
local timer_at = ngx.timer.at
local plugin_name = "http-logger-gelf"
local stale_timer_running = false
local buffers = {}

local schema = {
    type = "object",
    properties = {
        host = { type = "string" },
        port = { type = "integer" },
        timeout = { type = "integer", minimum = 1, default = 3 },
        name = { type = "string", default = "gelf http logger" },
        max_retry_count = { type = "integer", minimum = 0, default = 0 },
        retry_delay = { type = "integer", minimum = 0, default = 1 },
        buffer_duration = { type = "integer", minimum = 1, default = 60 },
        inactive_timeout = { type = "integer", minimum = 1, default = 5 },
        include_req_body = { type = "boolean", default = false },
    },
    required = { "host", "port" }
}

local _M = {
    version = 0.1,
    priority = 410,
    name = plugin_name,
    schema = schema,
}

function _M.check_schema(conf, schema_type)
    return core.schema.check(schema, conf)
end

local function send_http_data(conf, log_message)
    local err_msg
    local res = true

    local httpc = http.new()
    httpc:set_timeout(conf.timeout * 1000)
    local ok, err = httpc:connect(conf.host, conf.port)

    if not ok then
        return false, "failed to connect to host[" .. host .. "] port[" .. tostring(port) .. "] " .. err
    end

    local httpc_res, httpc_err = httpc:request({
        method = "POST",
        path = "/gelf",
        body = log_message,
        headers = {
            ["Content-Type"] = "application/json",
        }
    })

    if not httpc_res then
        return false, "error while sending data to [" .. host .. "] port[" .. tostring(port) .. "] " .. httpc_err
    end

    -- some error occurred in the server
    if httpc_res.status >= 400 then
        res = false
        err_msg = "server returned status code[" .. httpc_res.status .. "] host[" .. host .. "] port[" .. tostring(port) .. "] " .. "body[" .. httpc_res:read_body() .. "]"
    end

    return res, err_msg
end


-- remove stale objects from the memory after timer expires
local function remove_stale_objects(premature)
    if premature then
        return
    end

    for key, batch in ipairs(buffers) do
        if #batch.entry_buffer.entries == 0 and #batch.batch_to_process == 0 then
            core.log.warn("removing batch processor stale object, conf: ",
                    core.json.delay_encode(key))
            buffers[key] = nil
        end
    end

    stale_timer_running = false
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

    if not entry.route_id then
        entry.route_id = "no-matched"
    end

    if not stale_timer_running then
        -- run the timer every 30 mins if any log is present
        timer_at(1800, remove_stale_objects)
        stale_timer_running = true
    end

    local log_buffer = buffers[conf]

    if log_buffer then
        log_buffer:push(entry)
        return
    end

    -- Generate a function to be executed by the batch processor
    local func = function(entries, batch_max_size)
        local data, err = core.json.encode(entries[1]) -- encode as single {}

        if not data then
            return false, 'error occurred while encoding the data: ' .. err
        end

        return send_http_data(conf, data)
    end

    local config = {
        name = conf.name,
        retry_delay = conf.retry_delay,
        batch_max_size = conf.batch_max_size,
        max_retry_count = conf.max_retry_count,
        buffer_duration = conf.buffer_duration,
        inactive_timeout = conf.inactive_timeout,
        route_id = ctx.var.route_id,
        server_addr = ctx.var.server_addr,
    }

    local err
    log_buffer, err = batch_processor:new(func, config)

    if not log_buffer then
        core.log.error("error when creating the batch processor: ", err)
        return
    end

    buffers[conf] = log_buffer
    log_buffer:push(entry)
end

return _M