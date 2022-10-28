package main

import (
	"mall/middlewares"
	"mall/router"
)

func main() {
	r := router.Router()
	r.Use(middlewares.Cors())
	r.Run("") // 监听并在 0.0.0.0:8080 上启动服务
}
