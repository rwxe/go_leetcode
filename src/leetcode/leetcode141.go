package leetcode

func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}
	slowPointer := head
	fastPointer := head.Next
	for slowPointer != fastPointer {
		if fastPointer == nil || fastPointer.Next == nil {
			return false
		}
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next

	}
	return true

}
