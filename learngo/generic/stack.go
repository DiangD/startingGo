package main

type stack[T any] struct {
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

func main() {
	var strStack stack[string]
	strStack.push("qzh")
	strStack.push("test")
	strStack.push("demo")

	println(strStack.pop())
	println(strStack.size())
}
