package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

//gtfield greater than field = ?
type Booking struct {
	Checkin  time.Time `json:"checkin" form:"check_in" binding:"required,bookableDate" time_format:"2006-01-02"`
	Checkout time.Time `json:"checkout" form:"check_out" binding:"required,gtfield=Checkin" time_format:"2006-01-02"`
}

//v10版本的写法
func bookableDate(level validator.FieldLevel) bool {
	//类型断言
	if date, ok := level.Field().Interface().(time.Time); ok {
		today := time.Now()
		if date.Unix() > today.Unix() {
			return true
		}
	}
	return false
}

func main() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册验证器
		_ = v.RegisterValidation("bookableDate", bookableDate)
	}
	router.GET("/bookable", func(c *gin.Context) {
		var bookable Booking
		//绑定get param
		if err := c.ShouldBindQuery(&bookable); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "ok", "booking": bookable})
	})

	_ = router.Run(":8080")
}
