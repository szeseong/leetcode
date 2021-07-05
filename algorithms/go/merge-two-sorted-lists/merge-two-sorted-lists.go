package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	//var head *ListNode // nil pointer with ListNode type
	// if l1.Val < l2.Val {
	// 	head = l1
	// 	l1 = l1.Next
	// } else {
	// 	head = l2
	// 	l2 = l2.Next
	// }

	head := &ListNode{}
	tail := head
	for {
		if l1 == nil {
			tail.Next = l2
			break
		}

		if l2 == nil {
			tail.Next = l1
			break
		}

		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}

		tail = tail.Next
	}

	return head.Next
}
