package leetcode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	oldToNew := make(map[*Node]*Node)
	oldP := head
	newHead := &Node{Val: -1}
	newP := newHead
	for oldP != nil {
		newP.Next = &Node{
			Val:    oldP.Val,
			Next:   oldP.Next,
			Random: oldP.Random,
		}
		oldToNew[oldP] = newP.Next

		oldP = oldP.Next
		newP = newP.Next
	}
	newP = newHead.Next
	oldP = head
	for oldP != nil {
		if newP.Random != nil {
			newP.Random = oldToNew[oldP.Random]
		}
		oldP = oldP.Next
		newP = newP.Next
	}
	return newHead.Next
}
