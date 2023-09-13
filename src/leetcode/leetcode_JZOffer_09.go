package leetcode

type stackJZO09 []int

func (s *stackJZO09) push(x int) {
	*s = append(*s, x)
}
func (s *stackJZO09) pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
func (s *stackJZO09) empty() bool {
	return len(*s) == 0
}

type CQueueJZO09 struct {
	in, out stackJZO09
}

func ConstructorJZO09() CQueueJZO09 {
	return CQueueJZO09{
		in:  make(stackJZO09, 0),
		out: make(stackJZO09, 0),
	}
}

func (q *CQueueJZO09) AppendTail(value int) {
	q.in.push(value)
}

func (q *CQueueJZO09) DeleteHead() int {
	if q.in.empty() && q.out.empty() {
		return -1
	} else if q.out.empty() {
		for !q.in.empty() {
			q.out.push(q.in.pop())
		}
	}
	return q.out.pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
