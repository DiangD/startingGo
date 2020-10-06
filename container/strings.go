package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱一事无成的伟大" //UTF-8

	//和预想不一样是因为len获取的是字节数
	//可以用utf8.RuneCountInString(s)来获取字符串的长度
	fmt.Println(len(s))

	fmt.Println(s)
	//打印每一个字符的二进制
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	//utf8下英文为1字节 中文为3到4字节可变
	//utf8转化成Unicode
	for i, ch := range s {
		fmt.Printf("(%d,%X)", i, ch)
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()
}
