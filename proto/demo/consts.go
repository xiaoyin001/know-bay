package main

// 生成 .Lua 文件模板
const uLuaTemplate = `
{{- $PkgName := .PkgName -}}
{{- $SvrName := .svrData.Name -}}

-- 自动生成的 Lua gRPC 接口代码

local pb = require "pb"

{{- range .svrData.Comment -}}
-- {{ . }}
{{ end }}
local {{ $SvrName }} = {}

{{- range .svrData.RPCs }}
{{ range .Comment }}
-- {{ . }}
{{- end }}
{{- range .InlineComment }}
-- {{ . }}
{{- end }}
function {{ $SvrName }}.{{ .Name }} (req, callback)
    local function cb{{ .Name }} (code, response)
        if code ~= 0 then
            callback(code, response)
            return
        end

        local decodeMsg, err = pb.decode("{{ .ResponseType }}", response)
        if not decodeMsg then
            callback(-1, "{{ .ResponseType }} 解码失败: " .. err)
            return
        end

        callback(0, decodeMsg)
    end

    local encodeMsg = pb.encode("{{ .RequestType }}", req)

    Backend.AsyncRPC("", "/{{ $PkgName }}.{{ $SvrName }}/{{ .Name }}", encodeMsg, cb{{ .Name }})
end

{{- end }}

return {{ $SvrName }}
`
