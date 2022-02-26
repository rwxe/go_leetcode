package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func MergeNodes(head *ListNode) *ListNode {
	prev, p := head, head
	tempSum := 0

	for p != nil {
		if p.Val != 0 || p == prev {
			tempSum += p.Val
		} else {
			p.Val = tempSum
			tempSum = 0
			prev.Next = p
			prev = p
		}
		p = p.Next
	}

	return head.Next
}
