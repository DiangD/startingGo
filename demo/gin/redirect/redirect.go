package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	//重定向到外部链接
	router.GET("/to_blog", func(c *gin.Context) {
		c.Redirect(301, "https://shmiloveu.fun")
	})

	//重定向到内部链接
	router.GET("/to_test", func(c *gin.Context) {
		c.Request.URL.Path = "/test"
		router.HandleContext(c)
	})

	router.GET("/test", test)

	_ = router.Run()

}

//handleFunc
func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "redirect from /to_test",
	})
}
