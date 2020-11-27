package main

import (
	"fmt"
	"sync"
)

//.\escape_demo.go:34:6: can inline buildUserArr
//.\escape_demo.go:72:6: can inline printUser
//.\escape_demo.go:73:13: inlining call to fmt.Println
//.\escape_demo.go:48:6: can inline wrong.func1
//.\escape_demo.go:49:13: inlining call to printUser
//.\escape_demo.go:49:13: inlining call to fmt.Println
//.\escape_way.go:6:6: can inline pointerEscape
//.\escape_way.go:26:13: inlining call to fmt.Println
//.\escape_demo.go:55:21: inlining call to buildUserArr
//.\escape_demo.go:67:15: inlining call to pointerEscape
//.\escape_demo.go:61:13: inlining call to printUser
//.\escape_demo.go:61:13: inlining call to fmt.Println
//.\escape_demo.go:62:11: inlining call to sync.(*WaitGroup).Done
//.\escape_demo.go:36:15: []user literal escapes to heap
//.\escape_demo.go:46:12: leaking param content: arr
//.\escape_demo.go:48:11: leaking param: u
//.\escape_demo.go:47:9: moved to heap: u
//.\escape_demo.go:48:6: func literal escapes to heap
//.\escape_demo.go:49:13: []interface {} literal does not escape
//.\escape_way.go:18:11: make([]int, 10000, 10000) escapes to heap
//.\escape_way.go:26:13: s escapes to heap
//.\escape_way.go:26:13: []interface {} literal does not escape
//.\escape_demo.go:60:11: leaking param: u
//.\escape_demo.go:56:6: moved to heap: wg
//.\escape_demo.go:59:3: moved to heap: tmp
//.\escape_demo.go:55:21: []user literal does not escape
//.\escape_demo.go:60:6: func literal escapes to heap
//.\escape_demo.go:67:15: new(user) does not escape
//.\escape_demo.go:61:13: []interface {} literal does not escape
//.\escape_demo.go:72:16: leaking param: user
//.\escape_demo.go:73:13: []interface {} literal does not escape
//.\escape_way.go:8:10: new(user) escapes to heap

type user struct {
	username string
	password string
}

func buildUserArr() []user {
	//将实例返回，编译器会认为该实例会被继续使用，分配到堆
	return []user{
		{username: "qzh", password: "123"},
		{username: "gfh", password: "12345"},
		{username: "ff", password: "12345"},
		{username: "srw", password: "12345"},
	}
}

// u在调用过程中会逃逸到堆(在协程引用了别的协程的变量，变量逃逸到heap)，
//因此u不会随着函数结束而消亡
func wrong(arr []user) {
	for _, u := range arr {
		go func(u *user) {
			printUser(u)
		}(&u)
	}
}

func main() {
	arr := buildUserArr()
	var wg sync.WaitGroup
	wg.Add(len(arr))
	for _, u := range arr {
		tmp := u
		go func(u *user) {
			printUser(u)
			wg.Done()
		}(&tmp)
	}
	wg.Wait()

	pointerEscape()
	stackSizeExceed()
	dynamicType()
}

func printUser(user *user) {
	fmt.Println(user)
}
