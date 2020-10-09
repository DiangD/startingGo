package mock

type Retriever struct {
	Content string
}

//go没有显示的implement来实现接口，默认实现接口的所有方法即为实现接口
//通过这种形式来实现duck typing
func (r Retriever) Get(url string) string {
	return r.Content
}
