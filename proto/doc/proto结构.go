package main

import (
	"fmt"
	"os"

	"github.com/emicklei/proto"
)

/*

目前这些内容都是我目前使用到后的，后续有机会需要其他的部分也会进行继续补充的，有兴趣的可以继续关注

感谢proto库的作者，让我可以这么方便的解析.proto文件

*/

func main() {
	reader, _ := os.Open("test.proto")
	defer reader.Close()

	// 使用proto库创建一个解析器，用来解析 .proto 文件
	parser := proto.NewParser(reader)
	// 解析给定的Proto文件，获取其定义的内容
	definition, err := parser.Parse()
	if err != nil {
		panic(err)
	}

	// 遍历由解析器返回的定义
	proto.Walk(definition,
		proto.WithOption(handleOption),
		proto.WithPackage(handlePackage),
		proto.WithService(handleService),
		proto.WithMessage(handleMessage),
	)

	// 上面所有的数据都加载玩完了，然后再将这些数据按照Service为粒度的存文件
}

// 以Option为粒度，处理每个Option定义
func handleOption(o *proto.Option) {
	if o.Name != "lua_package" {
		return
	}

	if !o.Constant.IsString {
		return
	}

	if o.Constant.Source == "" {
		return
	}

	// 拿到.proto文件指定的Option
	fmt.Println(o.Constant.Source)
}

// 以Package为粒度，处理每个Package定义
func handlePackage(p *proto.Package) {
	fmt.Println("proto 的包名:", p.Name)
}

// 以Service为粒度，处理每个Service定义
func handleService(s *proto.Service) {

	// 查看接口的实际类型和值
	// fmt.Printf("实际类型和值: %T, %v\n", s.Elements, s.Elements)

	// Service 声明的位置【在.proto中的行列位置】
	fmt.Println("Service 声明的位置【在.proto中的行列位置】:", s.Position)
	// Service 的注释
	fmt.Println("Service 的注释:", s.Comment)
	// Service 的名字
	fmt.Println("Service 的名字:", s.Name)
	// Service 的构成部分
	fmt.Println("Service 的构成部分:", s.Elements)
	// Service 所有构成部分
	for i, serviceEach := range s.Elements {
		fmt.Printf("Service 构成部分的第%d个元素:%v\n", i, serviceEach)

		// 检查 each 是否为 *proto.RPC 类型，不是的直接跳过
		mRPC, ok := serviceEach.(*proto.RPC)
		if !ok {
			continue
		}

		// RPC 声明的位置
		fmt.Println("RPC 声明的位置:", mRPC.Position)
		// RPC 的注释
		fmt.Println("RPC 的注释:", mRPC.Comment)
		// RPC 的名字
		fmt.Println("RPC 的名字:", mRPC.Name)
		// RPC 的请求类型
		fmt.Println("RPC 的请求类型:", mRPC.RequestType)
		// RPC 的请求是否为streams
		fmt.Println("RPC 的请求是否为stream:", mRPC.StreamsRequest)
		// RPC 的响应类型
		fmt.Println("RPC 的响应类型:", mRPC.ReturnsType)
		// RPC 的响应是否为streams
		fmt.Println("RPC 的响应是否为stream:", mRPC.StreamsReturns)
		// RPC 的元素
		fmt.Println("RPC 的元素:", mRPC.Elements)
		// 遍历RPC的所有元素
		for j, rpcEach := range mRPC.Elements {
			// 目前看是没有的
			fmt.Printf("RPC 构成部分的第%d个元素:%v\n", j, rpcEach)
		}
		// RPC 的行内注释
		fmt.Println("RPC 的行内注释:", mRPC.InlineComment)
		// RPC 的父节点
		fmt.Println("RPC 的父节点:", mRPC.Parent) // mRPC.Parent.(*proto.Service)

	}
	// Service 的父节点
	fmt.Println("Service 的父节点:", s.Parent) // s.Parent.(*proto.Proto)

}

// 以Message为粒度，处理每个Message定义
func handleMessage(m *proto.Message) {

	// 结构体声明的位置
	fmt.Println("结构体声明的位置:", m.Position)
	// 结构体的注释
	fmt.Println("结构体的注释:", m.Comment)
	// 结构体的名字
	fmt.Println("结构体的名字:", m.Name)
	// 结构体是否为扩展
	fmt.Println("结构体是否为扩展:", m.IsExtend)
	// 结构体的元素
	fmt.Println("结构体的元素:", m.Elements)
	// 遍历Message的所有元素
	for i, messageEach := range m.Elements {
		fmt.Printf("Message 构成部分的第%d个元素:%v\n", i, messageEach)

		// 结构体下面的参数信息 messageEach.(*proto.NormalField)
		field, ok := messageEach.(*proto.NormalField)
		if !ok {
			continue
		}

		// 字段声明的位置
		fmt.Println("字段声明的位置:", field.Position)
		// 字段注释
		fmt.Println("字段注释:", field.Comment)
		// 字段名字
		fmt.Println("字段名字:", field.Name)
		// 字段类型
		fmt.Println("字段类型:", field.Type)
		// 字段序列
		fmt.Println("字段序列:", field.Sequence)
		// 字段的Options
		fmt.Println("字段的Options:", field.Options)
		// 字段行内注释
		fmt.Println("字段行内注释:", field.InlineComment)
		// 字段父节点
		fmt.Println("字段父节点:", field.Parent) // field.Parent.(*proto.Message)

		// 字段是否为重复【数组】
		fmt.Println("是否为数组:", field.Repeated)
		// 字段是否为可选字段
		fmt.Println("是否为可选字段:", field.Optional)
		// 字段是否为必须字段
		fmt.Println("是否为必须字段:", field.Required)
	}
	// 结构体的父节点
	fmt.Println("结构体的父节点:", m.Parent) // m.Parent.(*proto.Proto)
}
