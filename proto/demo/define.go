package main

// 存放 .proto 中 Service 字段
type protoService struct {
	Name    string     // serviceName
	Comment []string   // service 注释
	RPCs    []protoRPC // service 下的所有 rpc
}

// 存放 .proto 中 RPC 相关内容
type protoRPC struct {
	Name          string   // rpcName
	Comment       []string // rpc 上方注释
	InlineComment []string // rpc 同行注释
	RequestType   string   // rpc 请求类型
	ResponseType  string   // rpc 响应类型
}
