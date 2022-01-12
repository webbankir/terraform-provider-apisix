# terraform-provider-apisix

Only from 2.11.0 

Don't use without short hack in routes.lua

Patch method replace with 

```lua
function _M.patch(id, conf, sub_path, args)
    if not id then
        return 400, { error_msg = "missing route id" }
    end

    if conf == nil then
        return 400, { error_msg = "missing new configuration" }
    end

    if not sub_path or sub_path == "" then
        if type(conf) ~= "table" then
            return 400, { error_msg = "invalid configuration" }
        end
    end

    local key = "/routes"
    if id then
        key = key .. "/" .. id
    end

    local res_old, err = core.etcd.get(key)
    if not res_old then
        core.log.error("failed to get route [", key, "] in etcd: ", err)
        return 503, { error_msg = err }
    end

    if res_old.status ~= 200 then
        return res_old.status, res_old.body
    end
    core.log.info("key: ", key, " old value: ",
            core.json.delay_encode(res_old, true))

    local node_value = res_old.body.node.value
    local modified_index = res_old.body.node.modifiedIndex

    if sub_path and sub_path ~= "" then
        -- Main changes here
        if sub_path == "__full__" then
            node_value = conf
        else
            local code, err, node_val = core.table.patch(node_value, sub_path, conf)
            node_value = node_val
            if code then
                return code, err
            end
        end

        utils.inject_timestamp(node_value, nil, true)
    else
        node_value = core.table.merge(node_value, conf);
        utils.inject_timestamp(node_value, nil, conf)
    end

    core.log.info("new conf: ", core.json.delay_encode(node_value, true))

    local id, err = check_conf(id, node_value, true)
    if not id then
        return 400, err
    end

    local res, err = core.etcd.atomic_set(key, node_value, args.ttl, modified_index)
    if not res then
        core.log.error("failed to set new route[", key, "] to etcd: ", err)
        return 503, { error_msg = err }
    end

    return res.status, res.body
end
```