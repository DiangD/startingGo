package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Username string `json:"username" required:"true"`
	Password string `json:"password" required:"false"`
}

func main() {
	userType := reflect.TypeOf(user{})
	//遍历获取tag structTag string类型
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		fmt.Println(field.Tag)
	}
	if field, ok := userType.FieldByName("Password"); ok {
		fmt.Println(field.Tag.Get("json") + ":" + field.Tag.Get("required"))
	}

}
