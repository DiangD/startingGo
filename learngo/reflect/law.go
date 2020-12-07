package main

import (
	"fmt"
	"reflect"
)

//从接口到反射对象

func main() {
	user := "DiangD"

	//第一法则
	//反射的第一法则是我们能将 Go 语言的 interface{} 变量转换成反射对象。
	fmt.Println(reflect.TypeOf(user))
	//类型转换
	fmt.Println(reflect.ValueOf(user))

	//第二法则
	//反射的第二法则是我们可以从反射对象可以获取 interface{} 变量。
	v := reflect.ValueOf(1)
	//显式类型转换
	num := v.Interface().(int)
	fmt.Println("从反射对象可以获取 interface{} 变量:", num)

	i := 0
	//函数获取变量指针
	v = reflect.ValueOf(&i)
	//获取指针指向的变量，并设置新值
	v.Elem().SetInt(10)
	fmt.Println(i)
	//===========等价=============
	tmp := 10
	tmpVal := &tmp
	*tmpVal = 100
	fmt.Println("=======tmp=======", tmp)
}
