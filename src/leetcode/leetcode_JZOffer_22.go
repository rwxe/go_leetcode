package leetcode

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		k--
		if k < 0 {
			slow = slow.Next
		}

	}
	return slow
}
