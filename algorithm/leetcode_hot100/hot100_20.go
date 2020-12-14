package main

import "fmt"

//leetcode 热题 简单———括号匹配

/**
思路：辅助栈法
*/

func isValid(s string) bool {
	stack := make([]int32, 0)
	//记录当前栈顶的位置
	index := -1
	for _, ch := range s {
		if ch == '{' || ch == '[' || ch == '(' {
			//入栈
			stack = append(stack, ch)
			index++
		} else {
			if ch == '}' {
				if index < 0 || stack[index] != '{' {
					return false
				}
				//删除栈顶元素
				stack = stack[:index]
				index--

			}
			if ch == ']' {
				if index < 0 || stack[index] != '[' {
					return false
				}
				stack = stack[:index]
				index--

			}
			if ch == ')' {
				if index < 0 || stack[index] != '(' {
					return false
				}
				stack = stack[:index]
				index--
			}
		}
	}
	//栈是否为空
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
}
