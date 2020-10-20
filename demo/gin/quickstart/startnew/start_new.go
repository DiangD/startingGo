package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	//创建一个不包含中间件的路由
	r := gin.New()

	//全局中间件
	r.Use(gin.Logger())

	// 路由添加中间件，可以添加任意多个
	r.GET("/test", gin.Recovery(), testEndpoint)

	// 路由组中添加中间件
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
	authorized.Use(Authorized())
	{
		authorized.GET("/login", func(context *gin.Context) {
			context.String(200, "welcome to admin page!")
		})
	}
	_ = r.Run(":8080")
}

func Authorized() gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Query("username")
		log.Printf(context.FullPath()+"/username = %s", username)
		if username != "admin" {
			context.JSON(403, gin.H{
				"message": http.StatusText(403),
			})
			_ = context.Error(errors.New("你不是admin，没有权限"))
			//停止
			context.Abort()
		}
		//通过验证
		context.Next()
	}
}

func testEndpoint(context *gin.Context) {
	a, b := 1, 0
	_ = a / b
	context.String(200, "an error happened")
}
