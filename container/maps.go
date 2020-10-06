package main

import "fmt"

func main() {
	m := map[string]string{
		"name":   "DiangD",
		"course": "golang",
	}

	m2 := make(map[string]int)

	var m3 map[string]int
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Printf("key=%v,value=%v\n", k, v)
	}

	fmt.Println("Getting values")
	//第一个返回值为v，第二个返回值为是否存在
	course, ok := m["course"]
	fmt.Println(course, ok)
	//如果key不存在，返回的value为该类型的0值
	if exist, ok := m["exist"]; ok {
		fmt.Println(exist, ok)
	} else {
		fmt.Println("Key does not exist")
	}

	fmt.Println("Deleting values")
	//使用内建函数
	delete(m, "course")
	_, ok = m["course"]
	fmt.Println(ok)
}
