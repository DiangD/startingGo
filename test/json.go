package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type person struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	CreateAt string `json:"create_at"`
}

func main() {
	str := "{\n    \"id\":2333,\n    \"name\":\"qzh\",\n    \"age\":19,\n    \"create_at\":\"2006-09-12 11:22:01\"\n}"
	p := person{}
	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		fmt.Println(err.Error())
	}
	date, _ := time.ParseInLocation("2006-01-02 15:04:05", p.CreateAt, time.Local)
	fmt.Println(p)
	fmt.Println(date)
}
