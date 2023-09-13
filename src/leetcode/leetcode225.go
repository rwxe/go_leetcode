package leetcode

type queue232 []int

func (q *queue232) push(val int) {
	*q = append(*q, val)
}
func (q *queue232) top() int {
	x := (*q)[0]
	return x
}
func (q *queue232) pop() int {
	x := q.top()
	(*q) = (*q)[1:]
	return x
}
func (q *queue232) empty() bool {
	return len(*q) == 0
}

type MyStack225 struct {
	q1, q2 queue232
}

func Constructor225() MyStack225 {
	return MyStack225{
		q1: make(queue232, 0),
		q2: make(queue232, 0),
	}
}

func (s *MyStack225) Push(x int) {
	var emptyQ, nonemptyQ *queue232
	if s.q1.empty() {
		emptyQ, nonemptyQ = &s.q1, &s.q2
	} else {
		emptyQ, nonemptyQ = &s.q2, &s.q1
	}
	emptyQ.push(x)
	for !nonemptyQ.empty() {
		emptyQ.push(nonemptyQ.pop())
	}
}

func (s *MyStack225) Pop() int {
	var nonemptyQ *queue232
	if !s.q1.empty() {
		nonemptyQ = &s.q1
	} else {
		nonemptyQ = &s.q2
	}
	return nonemptyQ.pop()
}

func (s *MyStack225) Top() int {
	var nonemptyQ *queue232
	if !s.q1.empty() {
		nonemptyQ = &s.q1
	} else {
		nonemptyQ = &s.q2
	}
	return nonemptyQ.top()
}

func (s *MyStack225) Empty() bool {
	return s.q1.empty() && s.q2.empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
