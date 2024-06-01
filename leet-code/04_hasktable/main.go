package main

import "fmt"

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func lengthOfLongestSubstring(s string) int {
	right, left, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))
	for right < len(s) {
		if idx, ok := indexes[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		indexes[s[right]] = right
		right++
		res = max(res, right-left)
	}
	return res
}

// https://leetcode.cn/problems/lru-cache/description/
type Node struct {
	Key, Val   int
	Prev, Next *Node
}

type LRUCache struct {
	head, tail *Node
	Keys       map[int]*Node
	Cap        int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Keys: make(map[int]*Node), Cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node = &Node{Key: key, Val: value}
		this.Keys[key] = node
		this.Add(node)
	}

	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.Key)
		this.Remove(this.tail)
	}

}

func (this *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = this.head

	if this.head != nil {
		this.head.Prev = node
	}

	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.Next = nil
	}
}

func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		this.head = node.Next
		node.Next = nil
		return
	}

	if node == this.tail {
		this.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))

}
