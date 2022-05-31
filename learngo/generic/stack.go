package main

import "fmt"

type stack[T comparable] struct {
	values []T
}

func (s *stack[T]) push(element T) {
	s.values = append(s.values, element)
}

func (s *stack[T]) pop() (T, bool) {
	if len(s.values) == 0 {
		var null T
		return null, false
	}
	top := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return top, true
}

func (s *stack[T]) size() int {
	return len(s.values)
}

func (s *stack[T]) contains(element T) bool {
	for _, e := range s.values {
		if e == element {
			return true
		}
	}
	return false
}

func main() {
	var strStack stack[string]
	strStack.push("qzh")
	strStack.push("test")
	strStack.push("demo")

	fmt.Println(strStack.pop())
	fmt.Println(strStack.size())
	fmt.Println(strStack.contains("qzh"))
	fmt.Println(strStack.contains("啊对对对"))
}
