package leetcode

import (
	"fmt"
	"math"
	"strconv"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
   }

func RlistNode2Num(head *ListNode)int{
	stack:=make([]int,0)
	for head!=nil{
		stack=append(stack,head.Val)
		head=head.Next
	}
	num:=0
	for len(stack)>0{
		top:=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		num+=top*int(math.Pow(10,float64(len(stack))))
	}
	return num
}

func printRListNode(head *ListNode){
	for head!=nil{
		fmt.Printf("%d ",head.Val)
		head=head.Next
	}
	fmt.Println()
}

func num2RListNode(num int) *ListNode{
	//大数操作时会溢出
	numStr:=strconv.Itoa(num)
	head:= &ListNode{}
	p:=head
	for i:=len(numStr)-1;i>=0;i--{
		p.Val=int(numStr[i]-'0')
		if i!=0{
			p.Next=&ListNode{}
			p=p.Next
		}



	}
	return head
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head,tail *ListNode	
	carry:=0
	for l1!=nil || l2!=nil{
		num1,num2:=0,0
		if l1!=nil{
			num1=l1.Val
			l1=l1.Next
		}
		if l2!=nil{
			num2=l2.Val
			l2=l2.Next
		}
		sum:=num1+num2+carry
		sum,carry=sum%10,sum/10
		if head==nil{
			head=&ListNode{}
			tail=head
			tail.Val=sum
		}else{
			tail.Next=&ListNode{}
			tail=tail.Next
			tail.Val=sum
		}
	}
	if carry>0{
		tail.Next=&ListNode{Val:carry}
	}
	return head

}
