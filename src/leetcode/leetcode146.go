package leetcode


type LRUDoubleListNode struct {
	prev, next *LRUDoubleListNode
	Key, Value int
}

type LRUDoubleList struct {
	Head, Tail *LRUDoubleListNode
}

func NewLRUDoubleList() LRUDoubleList {
	l := LRUDoubleList{}
	l.Head = &LRUDoubleListNode{}
	l.Tail = &LRUDoubleListNode{}
	l.Head.next = l.Tail
	l.Tail.prev = l.Head
	return l
}

func (l *LRUDoubleList) DebugPrint() []int {
	result := make([]int, 0)
	p := l.Head
	for p != nil {
		result = append(result, p.Value)
		p = p.next
	}
	return result
}

func (l *LRUDoubleList) AddToHead(node *LRUDoubleListNode) {
	headNext := l.Head.next
	l.Head.next = node
	headNext.prev = node
	node.next = headNext
	node.prev = l.Head
}
func (l *LRUDoubleList) Remove(node *LRUDoubleListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (l *LRUDoubleList) RemoveTail() *LRUDoubleListNode {
	if l.Head.next == l.Tail {
		return nil
	}
	node := l.Tail.prev
	l.Remove(node)
	return node
}

type LRUCache struct {
	capacity  int
	keymap    map[int]*LRUDoubleListNode
	cacheList LRUDoubleList
}

func NewLRUCache(capacity int) LRUCache {
	lru := LRUCache{capacity: capacity}
	lru.keymap = make(map[int]*LRUDoubleListNode)
	lru.cacheList = NewLRUDoubleList()
	return lru
}
func (lru *LRUCache) makeRecently(key int) {
	node := lru.keymap[key]
	lru.cacheList.Remove(node)
	lru.cacheList.AddToHead(node)
}
func (lru *LRUCache) addItem(key, Value int) {
	node := &LRUDoubleListNode{Key: key, Value: Value}
	lru.keymap[key] = node
	lru.cacheList.AddToHead(node)
}
func (lru *LRUCache) removeItem(key int) {
	node := lru.keymap[key]
	lru.cacheList.Remove(node)
	delete(lru.keymap, key)
}
func (lru *LRUCache) removeLeast() {
	node := lru.cacheList.RemoveTail()
	delete(lru.keymap, node.Key)
}

func Constructor146(capacity int) LRUCache {
	return NewLRUCache(capacity)
}

func (lru *LRUCache) Get(key int) int {
	if item, ok := lru.keymap[key]; ok {
		lru.makeRecently(item.Key)
		return item.Value
	} else {
		return -1
	}

}

//
func (lru *LRUCache) Put(key int, value int) {
	if item, ok := lru.keymap[key]; ok {
		item.Value = value
		lru.makeRecently(item.Key)
	} else {
		if len(lru.keymap) >= lru.capacity {
			lru.removeLeast()
		}
		lru.addItem(key, value)
	}
}

