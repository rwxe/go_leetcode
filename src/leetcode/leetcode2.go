package leetcode

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{-1, nil}
	var r *ListNode
	carry := 0
	for r = result; l1 != nil || l2 != nil; r = r.Next {
		var sum int
		if l1 != nil && l2 != nil {
			sum = (l1.Val + l2.Val + carry) % 10

			carry = (l1.Val + l2.Val + carry) / 10

			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			sum = (l1.Val + carry) % 10

			carry = (l1.Val + carry) / 10

			l1 = l1.Next

		} else if l1 == nil && l2 != nil {
			sum = (l2.Val + carry) % 10

			carry = (l2.Val + carry) / 10

			l2 = l2.Next

		}
		sumNode := &ListNode{sum, nil}
		r.Next = sumNode
	}

	if carry != 0 {
		carryNode := &ListNode{1, nil}
		r.Next = carryNode
	}
	return result.Next
}
