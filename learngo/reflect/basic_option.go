package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age" default:"18"`
	Addr string `json:"addr"`
}

func (u user) PrintUser(in string) (string, int) {
	fmt.Printf("%s Name is %s, Age is %d \n", in, u.Name, u.Age)
	return u.Name, u.Age
}

type stringTmp string

func main() {
	u := user{"tom", 27, "beijing"}

	// 获取对象的 Value
	v := reflect.ValueOf(u)
	fmt.Println("Value:", v)
	// fmt.Printf("%v\n", u)

	// 获取对象的 Type
	t := reflect.TypeOf(u)
	fmt.Println("Type:", t)
	// fmt.Printf("%T\n", u)

	t1 := v.Type()
	fmt.Println(t == t1)

	//通过
	v1 := reflect.New(t)
	if v1.CanSet() {
		v1.SetInt(10)
	}
	fmt.Println(v1)
	fmt.Println()

	//kind 获取底层的数据类型
	tmp := stringTmp("123")
	//stringTmp
	fmt.Println(reflect.TypeOf(tmp))
	//string
	fmt.Println(reflect.TypeOf(tmp).Kind())

	// 修改反射对象的值
	i := 20
	fmt.Println("before i =", i)
	e := reflect.Indirect(reflect.ValueOf(&i))
	// e := reflect.ValueOf(&i).Elem()
	if e.CanSet() {
		e.SetInt(22)
	}
	fmt.Println("after i =", i)

	// 反射字段操作
	elem := reflect.ValueOf(&u).Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("field:%+v\n", field)
		tag := field.Tag
		fmt.Println("Tag:", tag)
		fmt.Println("Tag json:", tag.Get("json"))

		value := elem.Field(i)
		if value.CanSet() {
			if value.Kind() == reflect.Int {
				fmt.Println("change age to 30")
				value.SetInt(30)
			}
			if value.Kind() == reflect.String {
				fmt.Println("change name to jerry")
				value.SetString("jerry")
			}
		}
		fmt.Println("======================================")
	}
	fmt.Println("after user:", u)

	// go的反射无法访问私有成员与私有方法
	for i := 0; i < v.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Println("method name :", method.Name)
		mt := method.Type
		var in []reflect.Value
		for j := 0; j < mt.NumIn(); j++ {
			fmt.Println("method in type:", mt.In(j))
			if mt.In(j).Kind() == reflect.String {
				in = append(in, reflect.ValueOf("welcome"))
			}
			// 方法 1 获取的方法信息对象会把方法的接受者也当着入参之一
			if mt.In(j).Name() == t.Name() {
				in = append(in, v)
			}
		}
		// 获取方法返回类型
		for j := 0; j < mt.NumOut(); j++ {
			fmt.Println("method out type:", mt.Out(j))
		}

		out := method.Func.Call(in)
		for _, o := range out {
			fmt.Println("out:", o)
		}
	}
}
