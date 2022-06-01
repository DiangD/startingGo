package main

import "fmt"

type BuiltInOrdered interface {
	~string | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func max[T BuiltInOrdered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type myInt int

func main() {
	fmt.Println(max(10, 100))

	fmt.Println(max(myInt(10), myInt(100)))
}
