package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers445(l1 *ListNode, l2 *ListNode) *ListNode {
	//反转链表
	reverseList := func(head *ListNode) *ListNode {
		prev, next := (*ListNode)(nil), (*ListNode)(nil)
		for head != nil {
			next = head.Next
			head.Next = prev
			prev = head
			head = next
		}
		return prev
	}
	//链表长度计算
	listLength := func(head *ListNode) int {
		var length int
		for length = 0; head != nil; length++ {
			head = head.Next
		}
		return length
	}

	//这里将l1视为长链
	if listLength(l1) < listLength(l2) {
		l1, l2 = l2, l1
	}
	rl1 := reverseList(l1)
	rl2 := reverseList(l2)
	//p1Tail预先指向长链尾，用于最后的进位
	p1, p2, p1Tail := rl1, rl2, l1
	carry := false
	for p1 != nil || p2 != nil || carry {
		if p1 != nil && p2 != nil {
			if carry {
				p1.Val += 1
				carry = false
			}
			num := p1.Val + p2.Val
			p1.Val = num % 10
			carry = num >= 10

		} else if p1 != nil {
			if carry {
				num := p1.Val + 1
				p1.Val = num % 10
				carry = num >= 10
			}
		} else {
			if carry {
				//因为此时p1和p2都为nil，
				//需要另外的指针p1Tail预先指向p1最后一个结点
				p1Tail.Next = &ListNode{Val: 1}
				carry = false
			}
		}
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}
	return reverseList(rl1)

}
