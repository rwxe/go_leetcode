package leetcode

type stack232 []int

func (s *stack232) push(val int) {
	*s = append(*s, val)
}
func (s *stack232) pop() int {
	poped:=(*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return poped
}
func (s *stack232) empty() bool {
	if len(*s)==0{
		return true
	}else{
		return false
	}
}

type MyQueue struct {
	a,b stack232
}

/** Initialize your data structure here. */
func Constructor232() MyQueue {
	return MyQueue{}

}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	for !this.a.empty(){
		this.b.push(this.a.pop())
	}
	this.a.push(x)
	for !this.b.empty(){
		this.a.push(this.b.pop())
	}
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	return this.a.pop()

}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.a[len(this.a)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.a.empty()

}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

