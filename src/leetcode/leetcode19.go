package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd_3(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	count := n
	for count > 0 {
		fast = fast.Next
		count--
	}
	if fast != nil {
		for fast.Next != nil {
			fast = fast.Next
			slow = slow.Next
		}
		slow.Next = slow.Next.Next

	} else {
		if n > 1 {
			return head.Next
		} else {
			return nil
		}
	}
	return head
}
func removeNthFromEnd_2(head *ListNode, n int) *ListNode {
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}
	pos := length - n - 1
	if pos != -1 {
		prev := head
		for pos > 0 {
			prev = prev.Next
			pos--
		}
		prev.Next = prev.Next.Next
	} else {
		if length > 1 {
			return head.Next
		} else {
			return nil
		}
	}
	return head
}
func removeNthFromEnd_1(head *ListNode, n int) *ListNode {
	nodes := make([]*ListNode, 0)
	curr := head
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}
	pos := len(nodes) - n - 1
	if pos != -1 {
		prev := nodes[pos]
		prev.Next = prev.Next.Next
	} else {
		if len(nodes) > 1 {
			return nodes[1]
		} else {
			return nil
		}
	}
	return head

}
