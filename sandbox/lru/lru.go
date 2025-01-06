package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int                      // 缓存容量
	cache    map[int]*list.Element    // 哈希表：键 -> 链表节点
	list     *list.List               // 双向链表
}

// 缓存中的数据结构
type Pair struct {
	key   int
	value int
}

// Constructor 初始化 LRUCache
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

// Get 获取键对应的值
func (lru *LRUCache) Get(key int) int {
	if elem, ok := lru.cache[key]; ok {
		// 将元素移动到链表头部
		lru.list.MoveToFront(elem)
		// 返回值
		return elem.Value.(Pair).value
	}
	// 如果未找到，返回 -1
	return -1
}

// Put 插入或更新键值对
func (lru *LRUCache) Put(key, value int) {
	if elem, ok := lru.cache[key]; ok {
		// 如果键已存在，更新值，并移动到链表头部
		elem.Value = Pair{key: key, value: value}
		lru.list.MoveToFront(elem)
	} else {
		// 如果键不存在，先检查容量是否已满
		if lru.list.Len() >= lru.capacity {
			// 容量已满，移除链表尾部元素
			backElem := lru.list.Back()
			if backElem != nil {
				// 删除对应的哈希表项
				delete(lru.cache, backElem.Value.(Pair).key)
				// 移除链表中的节点
				lru.list.Remove(backElem)
			}
		}
		// 插入新节点到链表头部
		newElem := lru.list.PushFront(Pair{key: key, value: value})
		// 在哈希表中记录新节点
		lru.cache[key] = newElem
	}
}

func main() {
	// 测试 LRUCache
	obj := Constructor(2)
	obj.Put(1, 1) // 缓存中存入 (1,1)
	obj.Put(2, 2) // 缓存中存入 (2,2)

	fmt.Println(obj.Get(1)) // 输出: 1

	obj.Put(3, 3)           // 插入 (3,3)，淘汰最久未使用的键 2
	fmt.Println(obj.Get(2)) // 输出: -1（键 2 被淘汰）

	obj.Put(4, 4)           // 插入 (4,4)，淘汰最久未使用的键 1
	fmt.Println(obj.Get(1)) // 输出: -1（键 1 被淘汰）
	fmt.Println(obj.Get(3)) // 输出: 3
	fmt.Println(obj.Get(4)) // 输出: 4
}
