package main

//链表交点

// hash O(m+n)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]int)
	for headA != nil {
		m[headA] = headA.Val
		headA = headA.Next
	}
	for headB != nil {
		if headB.Val == m[headB] {
			return headB
		}
		headB = headB.Next

	}
	return nil
}

func getIntersectionNodeRomantic(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil { // 如果第一次遍历到链表尾部，就指向另一个链表的头部，继续遍历，这样会抵消长度差。如果没有相交，因为遍历长度相等，最后会是 nil ==  nil
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}

func main() {

}
