package leetcode

type MyLinkedList707Node struct {
	prev *MyLinkedList707Node
	next *MyLinkedList707Node
	val  int
}
type MyLinkedList707 struct {
	head   *MyLinkedList707Node
	tail   *MyLinkedList707Node
	length int
}

func Constructor707() MyLinkedList707 {
	ll := MyLinkedList707{}
	ll.head = &MyLinkedList707Node{val: -1}
	ll.tail = &MyLinkedList707Node{val: -1}
	ll.head.next = ll.tail
	ll.tail.prev = ll.head
	ll.length = 0
	return ll
}
func (ll *MyLinkedList707) getNode(index int) *MyLinkedList707Node {
	for p, i := ll.head, -1; p != ll.tail; p, i = p.next, i+1 {
		if i == index {
			return p
		}
	}
	return nil
}
func (ll *MyLinkedList707) Get(index int) int {
	node := ll.getNode(index)
	if node != nil {
		return node.val
	}
	return -1
}
func (ll *MyLinkedList707) AddAtIndex(index int, val int) {
	// index < 0 头插
	// index 有效 node前插
	// index == 0 尾插
	// index > 0 不插
	node := &MyLinkedList707Node{
		val: val,
	}

	var prevNode, nextNode *MyLinkedList707Node

	if index > ll.length {
		return
	}

	if index < 0 {
		prevNode = ll.head
		nextNode = ll.head.next
	} else if index == ll.length {
		prevNode = ll.tail.prev
		nextNode = ll.tail
	} else {
		nextNode = ll.getNode(index)
		prevNode = nextNode.prev
	}
	node.next = nextNode
	node.prev = prevNode
	nextNode.prev = node
	prevNode.next = node
	ll.length++
}

func (ll *MyLinkedList707) AddAtHead(val int) {
	ll.AddAtIndex(-1, val)
}

func (ll *MyLinkedList707) AddAtTail(val int) {
	ll.AddAtIndex(ll.length, val)
}

func (ll *MyLinkedList707) DeleteAtIndex(index int) {
	node := ll.getNode(index)
	if node == nil {
		return
	}
	prevNode := node.prev
	nextNode := node.next
	prevNode.next = nextNode
	nextNode.prev = prevNode
	ll.length--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
