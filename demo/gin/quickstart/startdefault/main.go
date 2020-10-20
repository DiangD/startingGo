package main

import "github.com/gin-gonic/gin"

func main() {
	//默认启动方式，包含 Logger、Recovery 中间件
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin",
		})
	})
	_ = r.Run()
}
