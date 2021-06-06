package leetcode

func removeElements(head *ListNode, val int) *ListNode {
	prev := (*ListNode)(nil)
	curr := head
	for curr != nil {
		if curr.Val == val {
			if curr == head {
				head = curr.Next
				curr=curr.Next
			} else {
				prev.Next = curr.Next
				curr=curr.Next
			}
		} else {

			prev, curr = curr, curr.Next
		}
	}
	return head

}
