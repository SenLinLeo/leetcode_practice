package main

import "fmt"

// LRUCache
/** 146. LRU 缓存
https://leetcode.cn/problems/lru-cache/description/
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
**/
type LRUCache struct {
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
	}
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

// moveToHead 删除节点
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// removeNode 删除当前节点，衔接前后的节点
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// addToHead 每次插入链表头部的后面一个节点
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// removeTail 删除尾部节点
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

type entry struct {
	key, value int
}

type DLinkedNode struct {
	key   int
	value int
	prev  *DLinkedNode
	next  *DLinkedNode
}

func main() {
	lruCache := Constructor(2)

	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	if val := lruCache.Get(1); val != 1 {
		fmt.Printf("Expected value 1, got %d", val)
	}

	lruCache.Put(3, 3)
	if val := lruCache.Get(2); val != -1 {
		fmt.Printf("Expected value -1, got %d", val)
	}

	lruCache.Put(4, 4)
	if val := lruCache.Get(1); val != -1 {
		fmt.Printf("Expected value -1, got %d", val)
	}
	if val := lruCache.Get(3); val != 3 {
		fmt.Printf("Expected value 3, got %d", val)
	}
	if val := lruCache.Get(4); val != 4 {
		fmt.Printf("Expected value 4, got %d", val)
	}
}
