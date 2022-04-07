package leetcode

type CQueueJZO09 struct {
	stack1, stack2 []int
}

func ConstructorJZO09() CQueueJZO09 {
	return CQueueJZO09{}
}

func (this *CQueueJZO09) AppendTail(value int) {
	for len(this.stack1) != 0 {
		node := this.stack1[len(this.stack1)-1]
		this.stack1 = this.stack1[:len(this.stack1)-1]
		this.stack2 = append(this.stack2, node)
	}
	this.stack1 = append(this.stack1, value)
	for len(this.stack2) != 0 {
		node := this.stack2[len(this.stack2)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
		this.stack1 = append(this.stack1, node)
	}
}

func (this *CQueueJZO09) DeleteHead() int {
	if len(this.stack1) != 0 {
		node := this.stack1[len(this.stack1)-1]
		this.stack1 = this.stack1[:len(this.stack1)-1]
		return node
	} else {
		return -1
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
