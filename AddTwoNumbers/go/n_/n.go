package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	el1 := l1
	el2 := l2
	extra := 0

	tail := &ListNode{}
	head := tail

	for {

		s := tail.Val
		if el1 != nil {
			s = s + el1.Val
			el1 = el1.Next
		}
		if el2 != nil {
			s = s + el2.Val
			el2 = el2.Next
		}

		if s > 9 {
			tail.Val = s % 10
			extra = 1
		} else {
			tail.Val = s
		}

		if el1 == nil && el2 == nil {
			if extra > 0 {
				tmp := &ListNode{}
				tmp.Val = extra
				tail.Next = tmp
				tail = tmp
			}
			return head
		}

		tmp := &ListNode{}
		tmp.Val = extra
		tail.Next = tmp
		tail = tmp

		extra = 0
	}
}
