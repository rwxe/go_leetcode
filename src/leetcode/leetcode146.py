from typing import Dict,  Optional, Union


class DoubleListNode(object):
    '''
    存储的键值对双向链表节点
    Attributes:
        key: 缓存的键
        value: 缓存的值
        prev: 前指针
        node: 后指针
    '''

    def __init__(self, key: Optional[int] = None, value: Optional[int] = None):
        self.key: Optional[int] = key
        self.value: Optional[int] = value
        self.prev: Optional[DoubleListNode] = None
        self.next: Optional[DoubleListNode] = None


class DoubleList(object):
    '''
    存储的键值对双向链表

    Attributes:
        head: 头指针，不包含有用数据，即哑节点
        tail: 尾指针，不包含有用数据，即哑节点
    '''

    def __init__(self):
        self.head = DoubleListNode()
        self.tail = DoubleListNode()
        self.head.next = self.tail
        self.tail.prev = self.head

    def add_to_head(self, node: DoubleListNode):
        '''
        添加新节点到双向节点的头部
        Args:
            node: 新节点。这个节点应该在调用前创建完成并传入
        Returns:
            None
        Raise:
            AttributeError: 找不到头节点的后节点
        '''
        head_next = self.head.next
        if head_next is None:
            raise AttributeError("head_next is None")
        self.head.next = node
        head_next.prev = node
        node.next = head_next
        node.prev = self.head

    def remove(self, node: DoubleListNode):
        '''
        从链中删除一个节点
        Args:
            node: 要删除的节点。这个节点应该在调用前通过其他方法找到其实例并传入
        Returns:
            None
        Raise:
            AttributeError: 找不到node的前后节点
        '''
        if node.prev is None or node.next is None:
            raise AttributeError("node next or prev is None")
        node.prev.next = node.next
        node.next.prev = node.prev

    def remove_from_tail(self) -> Union[DoubleListNode, None]:
        '''
        删除链表尾的节点
        Args:
            None
        Returns:
            如果链表包含有用数据，则返回被删除的节点实例。否则返回None
        Raise:
            AttributeError: 找不到tail的前节点
        '''
        if self.head.next == self.tail:
            return None
        else:
            if self.tail.prev is None:
                raise AttributeError("tail prev is None")
            node: DoubleListNode = self.tail.prev
            self.remove(node)
            return node


class LRUCache(object):
    '''
    LRU 缓存类，
    Attributes:
        cache: 缓存链表，每个节点存储了该缓存的键值对
        capacity: 容量，超过此容量将会抛弃最久未使用的缓存节点
        keymap: 哈希表，每个键对应一个缓存节点
    '''

    def __init__(self, capacity):
        '''
        初始化，指定容量
        '''
        self.capacity = capacity
        self.keymap: Dict[int, DoubleListNode] = {}
        self.cache = DoubleList()

    def add_item(self, key: int, value: int):
        '''
        添加新缓存
        Args:
            key: 缓存的键
            value: 缓存的值
        Returns:
            None
        Raise:
            None
        '''
        node = DoubleListNode(key=key, value=value)
        self.keymap[key] = node
        self.cache.add_to_head(node)

    def remove_item(self, key: int):
        '''
        删除缓存
        Args:
            key: 缓存的键
        Returns:
            None
        Raise:
            None
        '''
        node = self.keymap[key]
        self.cache.remove(node)
        del self.keymap[key]

    def remove_least_recently(self):
        '''
        删除最久未使用缓存
        Args:
            None
        Returns:
            None
        Raise:
            None
        '''
        node = self.cache.remove_from_tail()
        if node is None or node.key is None:
            return
        del self.keymap[node.key]

    def set_most_recently(self, key: int):
        '''
        设定某个键对应的缓存为最近使用
        Args:
            key: 缓存的键
        Returns:
            None
        Raise:
            None
        '''
        node = self.keymap[key]
        self.cache.remove(node)
        self.cache.add_to_head(node)

    def get(self, key: int) -> int:
        '''
        读取缓存
        Args:
            key: 缓存的键
        Returns:
            若存在该缓存，则返回该缓存的值，否则，返回-1
        Raise:
            None
        '''
        if key in self.keymap:
            value = self.keymap[key].value
            if value is not None:
                self.set_most_recently(key)
                return value
        return -1

    def put(self, key: int, value: int):
        '''
        新增缓存
        Args:
            key: 缓存的键
            value: 缓存的值
        Returns:
            None
        Raise:
            None
        '''
        if key in self.keymap:
            self.keymap[key].value = value
            self.set_most_recently(key)
        else:
            if len(self.keymap) >= self.capacity:
                self.remove_least_recently()
            self.add_item(key, value)


if __name__ == '__main__':
    lru = LRUCache(2)
    lru.put(1, 1)
    lru.put(2, 2)
    lru.get(1)
    lru.put(3, 3)
    lru.get(2)
    lru.put(4, 4)
    lru.get(1)
    lru.get(3)
    lru.get(4)
