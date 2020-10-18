package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(err)
		}
	}()
	a, b := 5, 0
	fmt.Println(a / b)
}

func main() {
	tryRecover()
}
