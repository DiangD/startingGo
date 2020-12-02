package main

import "fmt"

type person struct {
	sex  int
	name string
}

//类似于继承
type man struct {
	*person
	extra string
}

func (p person) print() {
	fmt.Println("this is a person")
}
func main() {
	m := man{}
	m.person = &person{
		sex:  0,
		name: "test",
	}
	m.extra = "extend person"
	fmt.Println(m.person, m.extra)
}
