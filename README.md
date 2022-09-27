# terraform-provider-apisix

!!!no longer supported by me !!!

Only from 2.11.0 

Don't use without next hook. [Docs](https://apisix.apache.org/docs/apisix/plugin-develop/#where-to-put-your-plugins) 


```lua
local apisix = require("apisix")
local core = require("apisix.core")

local old_http_init = apisix.http_init
apisix.http_init = function(...)

    local old_core_table_path = core.table.patch
    core.table.patch = function(node_value, sub_path, conf)
        if sub_path == "__patch_terraform_plugin_apisix__" then
            return old_core_table_path(conf, "", conf)
        else
            return old_core_table_path(node_value, sub_path, conf)
        end
    end

    old_http_init(...)
end
```
