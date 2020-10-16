package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	//能匹配/user/name 不能匹配/user
	router.GET("/user/:name", func(c *gin.Context) {
		res := "My name is %s"
		name := c.Param("name")
		res = fmt.Sprintf(res, name)
		c.String(http.StatusOK, res)
	})

	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	router.GET("/user/:name/*doing", func(c *gin.Context) {
		res := "My name is %s"
		name := c.Param("name")
		action := c.Param("doing")
		res = fmt.Sprintf(res, name)
		if action != "" {
			runes := []rune(action)
			res = res + " and I am " + string(runes[1:])
		}
		c.String(http.StatusOK, res)
	})

	//获取get请求的参数/admin?operator=?&time=?
	router.GET("/admin", func(c *gin.Context) {
		//可设置默认值
		t := c.DefaultQuery("time", time.Now().Format("2006-01-02 15:04:05"))
		operator := c.Query("operator")
		c.String(http.StatusOK, "%s is access to admin at %s", operator, t)
	})

	//获取post请求参数
	router.POST("/form_post", func(c *gin.Context) {
		//可设置默认值
		name := c.PostForm("name")
		url := c.DefaultPostForm("url", c.FullPath())
		id := c.PostForm("id")
		c.JSON(http.StatusOK, gin.H{
			"id":   id,
			"name": name,
			"url":  url,
			"date": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	//设置绑定端口
	_ = router.Run(":8080")
}
