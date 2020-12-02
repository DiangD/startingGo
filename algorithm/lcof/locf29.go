package main

import (
	"fmt"
	"math"
)

/*
剑指offer29 顺时针打印矩阵
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])

	var (
		total       = rows * columns
		dict        = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
		res         = make([]int, 0, total)
		row, column = 0, 0
		index       = 0
	)
	for i := 0; i < total; i++ {
		res = append(res, matrix[row][column])
		matrix[row][column] = math.MaxInt32
		nextRow, nextColumn := row+dict[index][0], column+dict[index][1]
		if nextRow < 0 || nextColumn < 0 || nextRow >= rows || nextColumn >= columns || matrix[nextRow][nextColumn] == math.MaxInt32 {
			index = (index + 1) % 4
		}
		row, column = row+dict[index][0], column+dict[index][1]
	}

	return res
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(spiralOrder(matrix))
}
