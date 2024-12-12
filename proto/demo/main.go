package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var uRootPath = flag.String("p", "proto", ".proto 文件所在的根目录")

func main() {
	flag.Parse()

	err := filepath.Walk(*uRootPath, walkFile)
	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", *uRootPath, err)
	}
}

// 处理 .proto 文件
func walkFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	if !strings.HasSuffix(path, ".proto") {
		return nil
	}

	handleFile(path)

	return nil
}
