package main

import (
	"embed"
	"fmt"
)

/**
Patterns must not match files outside the package's module, such as ‘.git/*’
embed无法匹配包模块以外的文件
*/

//go:embed test_embed/*
var fs embed.FS

//go:embed hello.txt
var s string

func main() {
	file, err := fs.ReadFile("test_embed/index.html")
	if err != nil {
		return
	}
	fmt.Println(string(file))
	fmt.Println("================================")
	fmt.Println(s)
}
