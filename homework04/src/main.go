package main

import (
	"my-go-project/src/pkg"
	"my-go-project/src/routes"
)

func main() {
	// 1. 初始化日志（必须最先初始化，后续模块依赖它）
	pkg.InitLogger()

	// 2. 初始化数据库
	pkg.InitDB()

	// 3. 初始化路由
	r := routes.SetupRouter()

	// 4. 启动服务
	r.Run(":8080")
}
