package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 定义了一个处理 /hello 路径的 GET 请求的路由。每当有客户端以 GET 方法请求 /hello 路径时，Gin 就会调用后面的匿名函数来处理这个请求。
	// c *gin.Context 是 Gin 提供的一个参数，它包含了请求的所有信息以及用来发送响应的方法。
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		// 200 是 HTTP 状态码，表示请求成功。
		// gin.H 是一个快捷方式来创建 map[string]interface{} 类型的数据结构，用于构建 JSON 响应体。
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
