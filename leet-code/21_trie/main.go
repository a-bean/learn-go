package main

// 208 实现 Trie (前缀树) https://leetcode.cn/problems/implement-trie-prefix-tree/description/

type Trie struct {
	isWord   bool
	children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{isWord: false, children: make(map[rune]*Trie)}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, w := range word {
		next, ok := cur.children[w]
		if !ok {
			next = &Trie{isWord: false, children: make(map[rune]*Trie)}
			cur.children[w] = next
		}
		cur = next
	}
	cur.isWord = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, w := range word {
		next, ok := cur.children[w]
		if !ok {
			return false
		}
		cur = next
	}
	return cur.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, w := range prefix {
		next, ok := cur.children[w]
		if !ok {
			return false
		}
		cur = next
	}
	return true
}

// 212 https://leetcode.cn/problems/word-search-ii/

func main() {

	var a = "ab"
	println(len(a))

}
