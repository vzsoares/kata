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

	tail := &ListNode{}
	head := tail

	extra := 0

	for {

		s := tail.Val
		if l1 != nil {
			s = s + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			s = s + l2.Val
			l2 = l2.Next
		}

		tail.Val = s % 10
		extra = s / 10

		if l1 == nil && l2 == nil {
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
