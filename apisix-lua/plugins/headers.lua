local core = require("apisix.core")
local plugin_name = "headers"
local ngx = ngx
local pairs = pairs
local type = type

local schema = {
    type = "object",
    properties = {
        response = {
            description = "new headers for response",
            type = "object",
            minProperties = 1,
        },
        request = {
            description = "new headers for request",
            type = "object",
            minProperties = 1,
        },
        sts = {
            description = "HTTP Strict-Transport-Security",
            type = "object",
            properties = {
                max_age = {
                    description = "The time, in seconds, that the browser should remember that a site is only to be accessed using HTTPS.",
                    type = "integer",
                    default = 31536000

                },
                include_sub_domains = {
                    description = "If this optional parameter is specified, this rule applies to all of the site's subdomains as well.",
                    type = "boolean",
                    default = true
                },
                preload = {
                    description = "",
                    type = "boolean",
                    default = true
                },
            },
            required = { "max_age" }
        }
    },
    minProperties = 1,
}

local _M = {
    version = 0.1,
    priority = 12500,
    name = plugin_name,
    schema = schema,
}

function _M.check_schema(conf)
    local ok, err = core.schema.check(schema, conf)
    if not ok then
        return false, err
    end

    if conf.response then
        for field, value in pairs(conf.response) do
            if type(field) ~= 'string' then
                return false, 'invalid type as header field'
            end

            if type(value) ~= 'string' and type(value) ~= 'number' then
                return false, 'invalid type as header value'
            end

            if #field == 0 then
                return false, 'invalid field length in header'
            end
        end
    end

    if conf.request then
        for field, value in pairs(conf.request) do
            if type(field) ~= 'string' then
                return false, 'invalid type as header field'
            end

            if type(value) ~= 'string' and type(value) ~= 'number' then
                return false, 'invalid type as header value'
            end

            if #field == 0 then
                return false, 'invalid field length in header'
            end
        end
    end

    return true
end

do
    function _M.access(conf, ctx)

        if not conf.request then
            return
        end

        if not conf.headers_request_arr then
            conf.headers_request_arr = {}

            for field, value in pairs(conf.request) do
                core.table.insert_tail(conf.headers_request_arr, field, value)
            end
        end

        local field_cnt = #conf.headers_request_arr
        for i = 1, field_cnt, 2 do
            core.request.set_header(ctx, conf.headers_request_arr[i],
                    core.utils.resolve_var(conf.headers_request_arr[i + 1], ctx.var))
        end
    end

    function _M.header_filter(conf, ctx)
        if not conf.headers_response_arr then
            conf.headers_response_arr = {}
        end
        if conf.sts then
            local sts = "max-age=" .. conf.sts.max_age .. ";"
            if conf.sts.include_sub_domains then
                sts = sts .. " includeSubDomains;"
            end

            if conf.sts.preload then
                sts = sts .. " preload;"
            end

            core.table.insert_tail(conf.headers_response_arr, "Strict-Transport-Security", sts)
        end

        if conf.response then
            for field, value in pairs(conf.response) do
                core.table.insert_tail(conf.headers_response_arr, field, value)
            end
        end

        local field_cnt = #conf.headers_response_arr
        for i = 1, field_cnt, 2 do
            local val = core.utils.resolve_var(conf.headers_response_arr[i + 1], ctx.var)
            ngx.header[conf.headers_response_arr[i]] = val
        end

    end
end  -- do


return _M
