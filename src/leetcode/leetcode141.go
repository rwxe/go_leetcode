package leetcode

func hasCycle(head *ListNode) bool {
	slowPointer := head
	fastPointer := head
	started := false
	if head==nil{
		return false
	}
	for {
		if slowPointer.Next != nil {
			slowPointer = slowPointer.Next
		} else {
			return false
		}
		if fastPointer.Next != nil {
			if fastPointer.Next.Next != nil {
				fastPointer = fastPointer.Next.Next
			}else{
				return false
			}
		} else {
			return false
		}
		if slowPointer == fastPointer && started {
			return true
		}
		started = true
	}

}
