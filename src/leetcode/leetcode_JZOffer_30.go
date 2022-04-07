package leetcode

type MinStackJZO30 struct {
	stack, minStack []int
}

/** initialize your data structure here. */
func ConstructorJZO30() MinStackJZO30 {
	return MinStackJZO30{}
}

func (this *MinStackJZO30) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, x)
	} else {
		if x <= this.minStack[len(this.minStack)-1] {
			this.minStack = append(this.minStack, x)
		}
	}
}

func (this *MinStackJZO30) Pop() {
	if this.stack[len(this.stack)-1] <= this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStackJZO30) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStackJZO30) Min() int {
	return this.minStack[len(this.minStack)-1]

}

/**
 * Your MinStackJZO30 object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
