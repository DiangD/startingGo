package main

import "fmt"

type Number interface {
	int8 | int16 | int32 | int64 | float32 | float64
}

func newGenericFunc[age Number](myAge age) {
	val := float64(myAge)
	fmt.Println(val)
}

func bubbleSort[N Number](input []N) []N {
	n := len(input)
	for i := 0; i < n-1; i++ {
		for j := i; j < n; j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	return input
}

func main() {
	var a int64 = 18
	var b = 18.6
	newGenericFunc(a)
	newGenericFunc(b)

	list := []int32{7, 6, 5, 4, 3, 2, 1}
	floatList := []float32{7, 6.6, 5.5, 4.3, 4.2, 0.1}
	fmt.Println(bubbleSort(list))
	fmt.Println(bubbleSort(floatList))
}
