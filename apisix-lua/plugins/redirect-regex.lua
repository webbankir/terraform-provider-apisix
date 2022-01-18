local ngx = ngx
local re_sub = ngx.re.sub
local plugin_name = "redirect-regex"
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
                    pattern = {
                        description = "Regex for full request",
                        type = "string",
                    },
                    replace = {
                        description = "Regex for full request",
                        type = "string",
                    },

                    status_code = {
                        type = "integer",
                        default = 301
                    }
                },
                allOf = {
                    required = { "pattern", "replace" }
                }
            }
        }
    }
}

local _M = {
    version = 0.1,
    priority = 1000,
    name = plugin_name,
    schema = schema,
}

function _M.rewrite(conf, ctx)

    local fullURI = ngx.var.scheme .. "://" .. ctx.var.http_host .. ctx.var.uri
    if ctx.var.args then
        fullURI = fullURI .. "?" .. ctx.var.args
    end

    for _, variant in ipairs(conf.variants)
    do
        local uri, _, _ = re_sub(fullURI, variant.pattern, variant.replace, "jo")
        if uri and uri ~= fullURI then
            return ngx.redirect(uri, variant.status_code);
        end
    end
end

return _M