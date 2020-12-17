package main

//回文链表

//思路：o(n) 将链表转化为数组，再对数组做遍历
func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	n := len(arr)
	for i, v := range arr[:n/2] {
		if v != arr[n-1-i] {
			return false
		}
	}
	return true
}

func main() {

}
