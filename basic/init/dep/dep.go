package main

import (
	"fmt"
	"shmiloveu.fun/startingGo/basic/init/pack"
	"shmiloveu.fun/startingGo/basic/init/util"
)

func main() {
	fmt.Println(pack.Pack)
	fmt.Println(util.Seed)
}

//同一个包不同源文件的init函数执行顺序，golang spec没做说明，执行顺序是源文件名称的字典序。
