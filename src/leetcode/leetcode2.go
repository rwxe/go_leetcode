package leetcode

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var p1, p2, long, short *ListNode
	for p1, p2 = l1, l2; p1 != nil && p2 != nil; p1, p2 = p1.Next, p2.Next {
	}
	if p2 == nil {
		long, short = l1, l2
	} else {
		long, short = l2, l1
	}

	var carry, tempSum int
	for p1, p2 = long, short; p1 != nil; p1 = p1.Next {
		if p2 != nil {
			tempSum = p1.Val + p2.Val + carry
			p2 = p2.Next
		} else if carry != 0 {
			tempSum = p1.Val + carry
		} else { //优化，无进位且短链遍历完了直接跳出
			break
		}
		p1.Val = tempSum % 10
		carry = tempSum / 10
	}
	if carry != 0 {
		for p1 = long; p1.Next != nil; p1 = p1.Next {
		}
		p1.Next = &ListNode{Val: 1, Next: nil}
	}
	return long
}
