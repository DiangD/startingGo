package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	res := ""
	for ; n > 0; n = n / 2 {
		tmp := n % 2
		res = strconv.Itoa(tmp) + res
	}
	return res
}

//使用系统接口Reader
func printFile(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13))
	s := `qzh
   is the most rich boy 
	in the world
`
	file, _ := os.Open("abc.txt")
	printFile(file)
	printFile(strings.NewReader(s))
}
