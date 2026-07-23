package main

import (
	"my-go-project/src/pkg"
	"my-go-project/src/routes"
)

func main() {
	// 1. 初始化数据库
	pkg.InitDB()

	// 2. 初始化路由
	r := routes.SetupRouter()

	// 3. 启动服务
	r.Run(":8080")
}
