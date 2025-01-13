package main

import (
	"container/list"
	"fmt"
	"sort"
)

// 3: 无重复字符的最长子串 https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
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

// 146: LRU 缓存机制 https://leetcode.cn/problems/lru-cache/description/
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

// 460: LFU 缓存机制 https://leetcode.cn/problems/lfu-cache/description/
type LFUCache struct {
	nodes    map[int]*list.Element
	lists    map[int]*list.List
	capacity int
	min      int
}

type node struct {
	key       int
	value     int
	frequency int
}

func LFUCacheConstructor(capacity int) LFUCache {
	return LFUCache{nodes: make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		min:      0,
	}
}

func (this *LFUCache) Get(key int) int {
	value, ok := this.nodes[key]
	if !ok {
		return -1
	}
	currentNode := value.Value.(*node)
	this.lists[currentNode.frequency].Remove(value)
	currentNode.frequency++
	if _, ok := this.lists[currentNode.frequency]; !ok {
		this.lists[currentNode.frequency] = list.New()
	}
	newList := this.lists[currentNode.frequency]
	newNode := newList.PushBack(currentNode)
	this.nodes[key] = newNode
	if currentNode.frequency-1 == this.min && this.lists[currentNode.frequency-1].Len() == 0 {
		this.min++
	}
	return currentNode.value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if currentValue, ok := this.nodes[key]; ok {
		currentNode := currentValue.Value.(*node)
		currentNode.value = value
		this.Get(key)
		return
	}
	if this.capacity == len(this.nodes) {
		currentList := this.lists[this.min]
		frontNode := currentList.Front()
		delete(this.nodes, frontNode.Value.(*node).key)
		currentList.Remove(frontNode)
	}
	this.min = 1
	currentNode := &node{
		key:       key,
		value:     value,
		frequency: 1,
	}
	if _, ok := this.lists[1]; !ok {
		this.lists[1] = list.New()
	}
	newList := this.lists[1]
	newNode := newList.PushBack(currentNode)
	this.nodes[key] = newNode
}

// 49 : 字母异位词分组 https://leetcode.cn/problems/group-anagrams/
// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Output:
// [
//
//	["ate","eat","tea"],
//	["nat","tan"],
//	["bat"]
//
// ]
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	record, res := map[string][]string{}, [][]string{}
	for _, str := range strs {
		sByte := []rune(str)
		sort.Sort(sortRunes(sByte))
		sstrs := record[string(sByte)]
		sstrs = append(sstrs, str)
		record[string(sByte)] = sstrs
	}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

}
