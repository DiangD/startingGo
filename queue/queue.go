package queue

/**
扩充系统类型和自定义类型
1. 定义别名
2. 使用组合（结构体）
*/

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	// v.(type) 可以通过断言来限制传入的类型但在运行时才知道错误
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
