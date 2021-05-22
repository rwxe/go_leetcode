package leetcode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func ReverseList(head *ListNode) *ListNode{
	var prev  *ListNode
	curr:=head
	var next  *ListNode
	for curr!=nil{
		next=curr.Next
		curr.Next=prev
		prev=curr
		curr=next

	}
	return prev


}

func isPalindrome0(head *ListNode) bool {
	if head==nil{
		return true
	}
	result1:=make([]int,0)
	result2:=make([]int,0)
	p:=head
	for p!=nil{
		result1=append(result1,p.Val)
		p=p.Next
	}
	tail:=ReverseList(head)
	p=tail
	for p!=nil{
		result2=append(result2,p.Val)
		p=p.Next
	}
	for i:=0;i<len(result1);i++{
		if result1[i]!=result2[i]{
			return false
		}
	}
	return true
	

}
