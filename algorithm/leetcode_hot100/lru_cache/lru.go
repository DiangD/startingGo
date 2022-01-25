package main

import "fmt"

//leetcode LRU缓存

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

// DLinkedNode 双向链表
type DLinkedNode struct {
	key, val   int
	prev, next *DLinkedNode
}

//初始化链表节点
func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		val: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		capacity: capacity,
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	if _, ok := l.cache[key]; !ok {
		return -1
	}
	node := l.cache[key]
	//更新到头
	l.moveToHead(node)
	return node.val
}

func (l *LRUCache) Put(key, value int) {
	//key不存在
	if _, ok := l.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		l.cache[key] = node
		//更新到头
		l.addToHead(node)
		l.size++
		//超出容量，删除尾部节点（最少访问节点）
		if l.size > l.capacity {
			node = l.removeTail()
			delete(l.cache, node.key)
			l.size--
		}
	} else {
		//更新value
		node := l.cache[key]
		node.val = value
		l.moveToHead(node)
	}
}

func (l *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) moveToHead(node *DLinkedNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeTail() *DLinkedNode {
	node := l.tail.prev
	l.removeNode(node)
	return node
}

func (l *LRUCache) printList() {
	head := l.head.next
	for head != l.tail {
		fmt.Printf("(%d,%d)\n", head.key, head.val)
		head = head.next
	}
}

func main() {
	lruCache := Constructor(3)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)
	lruCache.Put(4, 4)

	fmt.Println(lruCache.Get(4))
	lruCache.printList()
}
