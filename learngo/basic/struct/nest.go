package main

import "fmt"

type iAnimal interface {
	eat()
}

type animal struct {
	name string
}

func newAnimal(name string) *animal {
	return &animal{name: name}
}

func (a *animal) eat() {
	fmt.Println(a.name + " is eating")
}

//内嵌结构体，可以与oop语言的继承结合理解
type cat struct {
	*animal
}

func newCat(name string) *cat {
	return &cat{animal: newAnimal(name)}
}

//方法重载
func (cat *cat) eat() {
	fmt.Printf("children %v is eating\n", cat.name)
}

func main() {
	cat := newCat("cat")
	cat.eat()
}
