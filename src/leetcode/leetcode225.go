package leetcode

type queue232 []int

func (q *queue232) push(val int) {
	*q = append(*q, val)
}
func (q *queue232) pop() int {
	poped := (*q)[0]
	(*q) = (*q)[1:]
	return poped
}
func (q *queue232) empty() bool {
	if len(*q) == 0 {
		return true
	} else {
		return false
	}
}

type MyStack struct {
	a,b queue232
}

/** Initialize your data structure here. */
func Constructor225() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	for !this.a.empty(){
		this.b.push(this.a.pop())
	}
	this.a.push(x)
	for !this.b.empty(){
		this.a.push(this.b.pop())
	}

}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.a.pop()

}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.a[0]

}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.a.empty()

}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
