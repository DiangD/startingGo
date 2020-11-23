package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3
var ss = "KKK"

//定义包内变量
var (
	bb = 4
	cc = true
)

func variable() {
	var a int
	var s string
	fmt.Printf("%d %q \n", a, s)
}

//go 初始化变量
func variableInitialValue() {
	//go可以类型推断
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

//go类型推断
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "Shorter"
	fmt.Println(a, b, c, s)
}

func constDeclare() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		java
		python

		golang
		javaScript
	)
	fmt.Println(cpp, java, python, golang, javaScript)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)

}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	//强制类型转换
	c := int(math.Sqrt(float64(a*a + b*b)))
	return c
}
func main() {
	fmt.Println("Hello World")
	variable()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc, ss)
	euler()
	triangle()
	constDeclare()
	enums()
}
