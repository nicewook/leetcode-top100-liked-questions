// 2. Add Two Numbers - https://leetcode.com/problems/add-two-numbers/
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain
// a single digit. Add the two numbers and return it as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example:

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(a *ListNode, b *ListNode) *ListNode {
	head := new(ListNode)
	tail := head
	var num int

	for a != nil || b != nil {
		num = num / 10 // carry
		if a != nil {
			num += a.Val
			a = a.Next
		}
		if b != nil {
			num += b.Val
			b = b.Next
		}
		tail.Next = &ListNode{num % 10, nil}
		tail = tail.Next
	}
	if num/10 == 1 {
		tail.Next = &ListNode{num / 10, nil}
	}
	return head.Next
}