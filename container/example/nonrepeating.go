package main

import "fmt"

/*
输出最长不重复子串的长度
*/
func lengthOfNomRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {

		if lastIndex, ok := lastOccurred[ch]; ok && lastIndex >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
func main() {
	fmt.Println(lengthOfNomRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNomRepeatingSubStr("一如一如"))
}
