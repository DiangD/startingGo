package main

import (
	"fmt"
	mock "shmiloveu.fun/startingGo/learngo/retriever/mockretriever"
	real2 "shmiloveu.fun/startingGo/learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func Download(r Retriever) string {
	return r.Get("https://shmiloveu.fun")
}

//组合接口
type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "https://shmiloveu.fun"

func session(r RetrieverPoster) string {
	r.Post(url, map[string]string{
		"contents": "another fake https://shmiloveu.fun",
	})
	return r.Get(url)
}

func Post(p Poster) {
	p.Post("https://shmiloveu.fun", map[string]string{
		"name":   "DiangD",
		"course": "Golang",
	})
}

func main() {
	var r Retriever
	r = &mock.Retriever{
		Content: "test mockRetriever retriever",
	}
	inspect(r)
	r = &real2.Retriever{
		UserAgent: "Mozilla 5.0",
		Timeout:   time.Minute,
	}
	//fmt.Println(Download(r))
	//fmt.Println(Download(r1))
	inspect(r)

	//type assertion 类型断言，断言失败会直接panic
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Content)
	} else {
		fmt.Println("not a mockRetriever type")
	}

	retriever := mock.Retriever{Content: url}
	fmt.Println("==========Try a session=========")
	fmt.Println(session(&retriever))
}

//接口变量包含实现者的类型(可以为结构，也可以为指针) +实现者的值内容/指针
func inspect(r Retriever) {
	fmt.Printf("%T,%v\n", r, r)
	//type switch
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Content:", v.Content)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

//interface{} 表示任何类型，任何类型都实现了空接口
