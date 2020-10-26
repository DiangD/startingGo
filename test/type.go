package main

import (
	"fmt"
	"reflect"
)

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(len(nums))
	target := make([]interface{}, 0, len(nums))
	for _, num := range nums {
		target = append(target, num)
	}
	fmt.Println(len(target))
	fmt.Printf("%T %v\n", target, target)

	fmt.Println(reflect.ValueOf(nums), reflect.TypeOf(nums).Kind())
}
