--
-- Licensed to the Apache Software Foundation (ASF) under one or more
-- contributor license agreements.  See the NOTICE file distributed with
-- this work for additional information regarding copyright ownership.
-- The ASF licenses this file to You under the Apache License, Version 2.0
-- (the "License"); you may not use this file except in compliance with
-- the License.  You may obtain a copy of the License at
--
--     http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.
--
local core = require("apisix.core")
local expr = require("resty.expr.v1")
local plugin_name = "multi-response-rewrite"
local ngx = ngx
local pairs = pairs
local ipairs = ipairs

local schema = {
    type = "object",
    properties = {
        variants = {
            type = "array",
            minItems = 1,
            items = {
                type = "object",
                properties = {
                    headers = {
                        description = "new headers for response",
                        type = "object",
                        minProperties = 1,
                    },
                    body = {
                        description = "new body for response",
                        type = "string",
                    },
                    body_base64 = {
                        description = "whether new body for response need base64 decode before return",
                        type = "boolean",
                        default = false,
                    },
                    status_code = {
                        description = "new status code for response",
                        type = "integer",
                        minimum = 200,
                        maximum = 598,
                    },
                    vars = {
                        type = "array",
                        minItems = 1,
                    },
                },
                allOf = {
                    required = { "vars"}
                }
            }
        }
    }
}

local _M = {
    version = 0.1,
    priority = 899,
    name = plugin_name,
    schema = schema,
}

local function find_variant(conf, ctx)

    for _, variant in ipairs(conf.variants)
    do
        local response_expr, _ = expr.new(variant.vars)
        local match_result = response_expr:eval(ctx.var)
        if match_result then
            return variant
        end
    end
    return nil
end

function _M.check_schema(conf)
    local ok, err = core.schema.check(schema, conf)
    if not ok then
        return false, err
    end
    --
    --if conf.headers then
    --    for field, value in pairs(conf.headers) do
    --        if type(field) ~= 'string' then
    --            return false, 'invalid type as header field'
    --        end
    --
    --        if type(value) ~= 'string' and type(value) ~= 'number' then
    --            return false, 'invalid type as header value'
    --        end
    --
    --        if #field == 0 then
    --            return false, 'invalid field length in header'
    --        end
    --    end
    --end
    --
    --if conf.body_base64 then
    --    if not conf.body or #conf.body == 0 then
    --        return false, 'invalid base64 content'
    --    end
    --    local body = ngx.decode_base64(conf.body)
    --    if not body then
    --        return false, 'invalid base64 content'
    --    end
    --end
    --
    --if conf.vars then
    --    local ok, err = expr.new(conf.vars)
    --    if not ok then
    --        return false, "failed to validate the 'vars' expression: " .. err
    --    end
    --end

    return true
end

do

    function _M.body_filter(conf, ctx)
        if not ctx.multi_response_rewrite_variant then
            return
        end

        if ctx.multi_response_rewrite_variant.body then

            if ctx.multi_response_rewrite_variant.body_base64 then
                ngx.arg[1] = ngx.decode_base64(ctx.multi_response_rewrite_variant.body)
            else
                ngx.arg[1] = ctx.multi_response_rewrite_variant.body
            end

            ngx.arg[2] = true
        end
    end

    function _M.header_filter(conf, ctx)
        ctx.multi_response_rewrite_variant = find_variant(conf, ctx)
        if not ctx.multi_response_rewrite_variant then
            return
        end

        if ctx.multi_response_rewrite_variant.status_code then
            ngx.status = ctx.multi_response_rewrite_variant.status_code
        end

        if ctx.multi_response_rewrite_variant.body then
            core.response.clear_header_as_body_modified()
        end

        if not ctx.multi_response_rewrite_variant.headers then
            return
        end

        --reform header from object into array, so can avoid use pairs, which is NYI
        if not ctx.multi_response_rewrite_variant.headers_arr then
            ctx.multi_response_rewrite_variant.headers_arr = {}

            for field, value in pairs(ctx.multi_response_rewrite_variant.headers) do
                core.table.insert_tail(ctx.multi_response_rewrite_variant.headers_arr, field, value)
            end
        end

        local field_cnt = #ctx.multi_response_rewrite_variant.headers_arr
        for i = 1, field_cnt, 2 do
            local val = core.utils.resolve_var(ctx.multi_response_rewrite_variant.headers_arr[i + 1], ctx.var)
            ngx.header[ctx.multi_response_rewrite_variant.headers_arr[i]] = val
        end
    end

end  -- do


return _M
