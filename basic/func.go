package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//返回多个参数
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsuported operation: " + op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

//函数式编程
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args"+"(%d,%d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//可变参数列表
func sum(nums ...int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}

func main() {
	if res, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {

		fmt.Println(res)
	}
	q, r := div(13, 4)
	fmt.Println(q, r)
	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))

}
