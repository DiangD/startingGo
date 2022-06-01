package main

import (
	"fmt"
	"reflect"
)

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func Convert[T1, T2 Integer](in T1) T2 {
	return T2(in)
}

func main() {
	var a int8 = 10
	b := Convert[int8, int64](a) // can't infer the return type
	fmt.Println(b, reflect.TypeOf(b))
}
