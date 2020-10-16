package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	//单文件上传
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		if file != nil {
			err := c.SaveUploadedFile(file, "/")
			if err != nil {
				//todo
			}
			c.JSON(200, gin.H{
				"fileName": file.Filename,
				"size":     file.Size,
			})
		}
	})

	//多文件上传
	router.POST("/uploads", func(c *gin.Context) {
		// 多文件
		form, _ := c.MultipartForm()
		files := form.File["files"]
		list := make([]string, 0)
		for _, file := range files {
			list = append(list, file.Filename)
			// 上传文件到指定的路径
			// c.SaveUploadedFile(file, dst)
		}
		c.JSON(200, gin.H{
			"files": list,
			"len":   len(list),
		})
	})
	_ = router.Run(":8080")
}
