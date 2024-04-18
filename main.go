package main

import (
	"echo-go/router"
	"fmt"
)

func main() {
	fmt.Println("hello")

	// 使用 := 短变量声明语法声明并初始化一个变量
	r := router.InitRouter()

	// 启动服务器
	err := r.Run(":8082")
	if err != nil {
		return
	}
}
