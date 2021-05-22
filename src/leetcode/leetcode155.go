package leetcode

type MinStack struct {
    stack []int
    minStack []int

}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{[]int{},[]int{1<<31}}

}


func (this *MinStack) Push(val int)  {
    this.stack=append(this.stack,val)
    if val<this.minStack[len(this.minStack)-1]{
        this.minStack=append(this.minStack,val)
    }else{
        this.minStack=append(this.minStack,this.minStack[len(this.minStack)-1])
	}


}


func (this *MinStack) Pop()  {
	this.stack=this.stack[:len(this.stack)-1]
	this.minStack=this.minStack[:len(this.minStack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]

}


func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]

}
