package main

import "fmt"

type printableInt interface {
	~int
	String()
}

type myInt int

func (m myInt) String() {
	fmt.Println(m)
}

type impossiblePrintInt interface {
	int
	String()
}

type ImpossibleStruct[T impossiblePrintInt] struct {
	val T
}

func (i ImpossibleStruct[T]) String() {
	fmt.Println(i)
}

func main() {
	s := ImpossibleStruct[int]{10}
	s.String()
	s2 := ImpossibleStruct[myInt]{10}
	s2.String()
}
