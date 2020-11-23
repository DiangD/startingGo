package main

import "fmt"

//函数式编程 闭包概念
//https://zh.wikipedia.org/wiki/%E9%97%AD%E5%8C%85_(%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%A7%91%E5%AD%A6)
func adder() func(int) int {
	sum := 0
	return func(num int) int {
		sum += num
		return sum
	}
}

type iAdder func(int) (int, iAdder)

//正统函数式编程 参数只能有一个，变量不能有状态
func adderOrigin(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adderOrigin(base + v)
	}

}

func main() {
	function := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(function(i))
	}
}
