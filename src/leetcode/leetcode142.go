package leetcode

func detectCycle(head *ListNode) *ListNode {
	if head==nil || head.Next==nil{
		return nil
	}
	slowPointer:=head
	fastPointer:=head
	notStarted:=true
	for slowPointer!=fastPointer || notStarted{
		notStarted=false
		if fastPointer==nil || fastPointer.Next==nil{
			return nil
		}
		fastPointer=fastPointer.Next.Next
		slowPointer=slowPointer.Next
	}
	p:=head
	for p!=slowPointer{
		p=p.Next
		slowPointer=slowPointer.Next
	}
	return 




}
