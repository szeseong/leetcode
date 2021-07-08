package main

import (
	"log"
	"testing"
)

func TestTwoEmptyListNode(t *testing.T) {
	var nil1Node *ListNode
	var nil2Node *ListNode
	got := mergeTwoLists(nil1Node, nil2Node)

	if got != nil {
		t.Errorf("%v", got)
	}

}

func TestTwoListNode(t *testing.T) {
	l4 := ListNode{Val: 10}
	l3 := ListNode{Val: 6, Next: &l4}
	//l3 := ListNode{Val: 6}
	l2 := ListNode{Val: 3, Next: &l3}
	l1 := ListNode{Val: 1, Next: &l2}

	//l1 := ListNode{}
	k3 := ListNode{Val: 10}
	k2 := ListNode{Val: 2, Next: &k3}
	k1 := ListNode{Val: 2, Next: &k2}

	got := mergeTwoLists(&l1, &k1)

	for ; got != nil; got = got.Next {
		log.Printf("loop %v\n", got.Val)
	}

	if got != nil {
		t.Errorf("%v", got)
	}

}
