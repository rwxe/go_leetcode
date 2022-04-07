package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newHead := &ListNode{Val: -1}
	p := newHead
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				p.Next = list1
				list1 = list1.Next
				p = p.Next
			} else {
				p.Next = list2
				list2 = list2.Next
				p = p.Next
			}
		} else if list2 == nil {
			p.Next = list1
			break
		} else {
			p.Next = list2
			break

		}
	}
	return newHead.Next
}
