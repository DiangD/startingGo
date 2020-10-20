package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	gin.DisableConsoleColor()

	file, _ := os.Create("demo/gin/log/writelog/gin.log")
	router := gin.Default()

	//只写入文件
	//gin.DefaultWriter = io.MultiWriter(file)

	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	router.GET("/log", func(context *gin.Context) {
		context.String(200, "writing log...")
	})
	_ = router.Run(":8080")
}
