package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Username string `json:"username" uri:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" uri:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	//绑定json
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.Username != "admin" || json.Password != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	//绑定表单
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		//默认绑定表单
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.Username != "user" || form.Password != "user" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	//绑定uri
	router.GET("/:username/:password", func(c *gin.Context) {
		var uri Login
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if uri.Username != "admin" || uri.Password != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	_ = router.Run(":8080")
}
