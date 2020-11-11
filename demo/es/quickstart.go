package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}
	//获取es版本
	version, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		panic(err)
	}
	fmt.Println("========elasticsearch======= version:", version)
	user := User{
		Name:    "DiangD",
		Age:     20,
		Married: false,
	}
	resp, err := client.Index().
		Index("user").
		BodyJson(user).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", resp.Id, resp.Index, resp.Type)
}
