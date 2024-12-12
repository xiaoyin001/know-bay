package main

import (
	"fmt"
	"os"
	"sync"
	"text/template"

	"github.com/emicklei/proto"
)

// 读取文件，进行处理
func handleFile(filePath string) {
	initCommonParam()

	reader, _ := os.Open(filePath)
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
	)

	// 上面所有的数据都加载玩完了，然后再将这些数据按照Service为粒度的存文件
	saveDataToLuaFile()
}

// 根据解析的 .Proto 文件，按照 Service 的粒度存入 .lua 文件
func saveDataToLuaFile() {
	fmt.Println("lua的存放路径:", uLuaPath)
	fmt.Println(".proto 中 PackageName:", uProtoPackageName)

	wg := sync.WaitGroup{}
	wg.Add(len(uProtoServices))

	for i, serrvice := range uProtoServices {
		svr := serrvice
		fmt.Printf("文件中有 %d 个 Service 第 %d 个 Service 的名字是 %s 备注是 %s\n", len(uProtoServices), i+1, svr.Name, svr.Comment)

		for j, rpc := range svr.RPCs {
			fmt.Printf("第 %d 个RPC的信息: %v\n", j+1, rpc)
		}

		// =====================================================================

		luaFilePath := fmt.Sprintf("%s%s.lua", uLuaPath, svr.Name)
		tmpl := template.Must(template.New("lua").Parse(uLuaTemplate))

		go func() {
			defer wg.Done()

			file, err := os.Create(luaFilePath)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			data := map[string]interface{}{
				"svrData": svr,
				"PkgName": uProtoPackageName,
			}
			err = tmpl.Execute(file, data)
			if err != nil {
				panic(err)
			}
		}()
	}

	wg.Wait()
}

// 检查文件夹是否存在，如果不存在则创建
func ensureDir(dirName string) error {
	info, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		return os.MkdirAll(dirName, 0o755)
	}
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return &os.PathError{Op: "mkdir", Path: dirName, Err: os.ErrExist}
	}

	return nil
}
