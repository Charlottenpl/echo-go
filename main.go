package main

import (
	"echo-go/router"
	"echo-go/sql"
	"fmt"
)

func main() {
	fmt.Println("hello")

	// init sql
	_, err := sql.Db()
	if err != nil {
		return
	}
	// 使用 := 短变量声明语法声明并初始化一个变量
	r := router.InitRouter()

	// 启动服务器
	err = r.Run(":8082")
	if err != nil {
		return
	}
}
