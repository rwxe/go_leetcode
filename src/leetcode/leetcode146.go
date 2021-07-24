package leetcode

// 双向链表中的节点
type LRUDoubleListNode struct {
	prev, next *LRUDoubleListNode
	key, value int
}

// 双向链表
type LRUDoubleList struct {
	head, tail *LRUDoubleListNode
}

// 构造一个初始双向链表
func NewLRUDoubleList() LRUDoubleList {
	l := LRUDoubleList{}
	l.head = &LRUDoubleListNode{}
	l.tail = &LRUDoubleListNode{}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

// 将节点添加到链表首
func (l *LRUDoubleList) addToHead(node *LRUDoubleListNode) {
	headNext := l.head.next
	l.head.next = node
	headNext.prev = node
	node.next = headNext
	node.prev = l.head
}

// 移除节点
func (l *LRUDoubleList) remove(node *LRUDoubleListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 移除链表尾节点
func (l *LRUDoubleList) removeFromTail() *LRUDoubleListNode {
	if l.head.next == l.tail {
		return nil
	} else {
		node := l.tail.prev
		l.remove(node)
		return node
	}
}

// 缓存结构
type LRUCache struct {
	capacity int
	keymap   map[int]*LRUDoubleListNode
	cache    LRUDoubleList
}

// 构造一个缓存
func NewLRUCache(capacity int) LRUCache {
	lru := LRUCache{capacity: capacity}
	lru.keymap = make(map[int]*LRUDoubleListNode)
	lru.cache = NewLRUDoubleList()
	return lru
}

// 设某个键为最近访问
func (lru *LRUCache) setMostRecently(key int) {
	node := lru.keymap[key]
	lru.cache.remove(node)
	lru.cache.addToHead(node)
}

// 新增元素
func (lru *LRUCache) addItem(key, Value int) {
	node := &LRUDoubleListNode{key: key, value: Value}
	lru.keymap[key] = node
	lru.cache.addToHead(node)
}

// 移除元素
func (lru *LRUCache) removeItem(key int) {
	node := lru.keymap[key]
	lru.cache.remove(node)
	delete(lru.keymap, key)
}

// 移除最久未访问元素
func (lru *LRUCache) removeLeastRecently() {
	node := lru.cache.removeFromTail()
	delete(lru.keymap, node.key)
}

// 读取操作
func (lru *LRUCache) Get(key int) int {
	if item, ok := lru.keymap[key]; ok {
		lru.setMostRecently(item.key)
		return item.value
	} else {
		return -1
	}

}

// 新增操作
func (lru *LRUCache) Put(key int, value int) {
	if item, ok := lru.keymap[key]; ok {
		item.value = value
		lru.setMostRecently(item.key)
	} else {
		if len(lru.keymap) >= lru.capacity {
			lru.removeLeastRecently()
		}
		lru.addItem(key, value)
	}
}

func Constructor146(capacity int) LRUCache {
	return NewLRUCache(capacity)
}
