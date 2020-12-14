package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

func createLinkedList(values []int) *ListNode {
	var head *ListNode = nil
	var curr *ListNode = nil
	for i, val := range values {
		if i == 0 {
			head = &ListNode{
				Val: val,
			}
			curr = head
		} else {
			if curr != nil {
				curr.Next = &ListNode{
					Val: val,
				}
				curr = curr.Next
			}
		}
	}
	return head
}

func printList(head *ListNode) {
	values := make([]string, 0)
	for head != nil {
		values = append(values, strconv.Itoa(head.Val))
		head = head.Next
	}
	fmt.Println(strings.Join(values, "->"))
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	head := createLinkedList(nums)
	printList(head)
	head = reverseList(head)
	printList(head)
}
