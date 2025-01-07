package hashmap

import (
	"fmt"
	"hash/fnv"
)

var defaultCapacity uint64 = 1 << 10

type node struct {
	key   any
	value any
	next  *node
}

// HashMap is golang implementation of hashmap
type HashMap struct {
	capacity uint64
	size     uint64
	table    []*node
}

// New return new HashMap instance
func New() *HashMap {
	return &HashMap{
		capacity: defaultCapacity,
		table:    make([]*node, defaultCapacity),
	}
}

// Make creates a new HashMap instance with input size and capacity
func Make(size, capacity uint64) HashMap {
	return HashMap{
		size:     size,
		capacity: capacity,
		table:    make([]*node, capacity),
	}
}

// Get returns value associated with given key
func (hm *HashMap) Get(key any) any {
	node := hm.getNodeByHash(hm.hash(key))

	if node != nil {
		return node.value
	}

	return nil
}

// Put puts new key value in hashmap
func (hm *HashMap) Put(key any, value any) any {
	return hm.putValue(hm.hash(key), key, value)
}

// Contains checks if given key is stored in hashmap
func (hm *HashMap) Contains(key any) bool {
	node := hm.getNodeByHash(hm.hash(key))
	return node != nil
}

func (hm *HashMap) putValue(hash uint64, key any, value any) any {
	// 如果哈希表还未初始化，进行初始化操作
	if hm.capacity == 0 {
		hm.capacity = defaultCapacity             // 设置默认容量（1024）
		hm.table = make([]*node, defaultCapacity) // 创建哈希表数组
	}

	// 根据哈希值获取对应位置的节点
	node := hm.getNodeByHash(hash)

	// 情况1：该位置还没有节点
	if node == nil {
		hm.table[hash] = newNode(key, value) // 直接创建新节点放入

		// 情况2：该位置已有节点，且key相同（更新值）
	} else if node.key == key {
		// 创建新节点，并将原节点作为next（形成链表）
		hm.table[hash] = newNodeWithNext(key, value, node)
		return value

		// 情况3：发生了哈希冲突（不同的key映射到相同位置）
	} else {
		hm.resize() // 扩容哈希表
		// 递归调用，重新插入该键值对
		return hm.putValue(hash, key, value)
	}

	// 插入新节点后，增加哈希表的大小计数
	hm.size++

	return value
}

func (hm *HashMap) getNodeByHash(hash uint64) *node {
	return hm.table[hash]
}

func (hm *HashMap) resize() {
	// 容量翻倍（左移1位相当于乘2）
	hm.capacity <<= 1

	// 保存旧的哈希表
	tempTable := hm.table

	// 创建新的、更大的哈希表
	hm.table = make([]*node, hm.capacity)

	// 遍历旧表中的所有位置
	for i := 0; i < len(tempTable); i++ {
		node := tempTable[i]
		// 跳过空位置
		if node == nil {
			continue
		}

		// 对于非空节点，使用新的容量重新计算哈希值
		// 并将节点放入新表中对应位置
		// 注意：由于容量变化，同一个key的哈希值会发生变化
		hm.table[hm.hash(node.key)] = node
	}
}

func newNode(key any, value any) *node {
	return &node{
		key:   key,
		value: value,
	}
}

func newNodeWithNext(key any, value any, next *node) *node {
	return &node{
		key:   key,
		value: value,
		next:  next,
	}
}

func (hm *HashMap) hash(key any) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))

	hashValue := h.Sum64()

	return (hm.capacity - 1) & (hashValue ^ (hashValue >> 16))
}
