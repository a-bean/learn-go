// 字典树（Trie）的数据结构，主要用于高效地存储和查找字符串。

/*
使用场景的例子：

1. 自动补全
在搜索引擎或文本输入框中，用户输入部分字符串时，可以使用字典树快速查找以该字符串为前缀的所有单词，从而提供自动补全建议。

2. 拼写检查
字典树可以用于拼写检查工具，通过存储正确的单词列表，快速判断用户输入的单词是否存在于字典中。

3. 前缀匹配
在需要查找以特定前缀开头的字符串时，字典树能够高效地返回所有匹配的字符串，例如在社交媒体平台中查找以某个标签开头的帖子。

4. IP 地址路由
在网络路由中，字典树可以用于存储和查找 IP 地址前缀，以便快速确定数据包的路由。

5. 词频统计
在文本分析中，可以使用字典树来统计每个单词的出现频率，适合处理大量文本数据。

6. 多语言支持
字典树可以存储多种语言的单词，支持多语言输入和查找，适合国际化应用。
*/
package trie

type ITrie interface {
	Insert(s ...string)
	Find(s string) bool
	Size() int
	Capacity() int
	Remove(s ...string)
	Compact() bool
}
type Node struct {
	children map[rune]*Node // 孩子节点的映射
	isWord   bool           // 代表一个完整的字符串。
}

// NewNode 创建一个新的节点
func NewNode() *Node {
	n := &Node{}
	n.children = make(map[rune]*Node) // 初始化孩子节点映射
	n.isWord = false                  // 默认不是叶子节点
	return n
}

// insert 插入一个字符串到当前节点
func (n *Node) insert(s string) {
	curr := n
	for _, c := range s {
		next, ok := curr.children[c]
		if !ok {
			next = NewNode() // 创建新节点
			curr.children[c] = next
		}
		curr = next // 移动到下一个节点
	}
	curr.isWord = true // 标记为叶子节点
}

// Insert 批量插入字符串
func (n *Node) Insert(s ...string) {
	for _, ss := range s {
		n.insert(ss) // 调用 insert 方法
	}
}

// Find 查找字符串是否存在
func (n *Node) Find(s string) bool {
	next, ok := n, false
	for _, c := range s {
		next, ok = next.children[c]
		if !ok {
			return false // 字符串不存在
		}
	}
	return next.isWord // 返回是否为叶子节点
}

// Capacity 返回树的容量 即节点的总数
func (n *Node) Capacity() int {
	r := 0
	for _, c := range n.children {
		r += c.Capacity() // 递归计算容量
	}
	return 1 + r // 当前节点加上所有孩子节点的容量
}

// Size 返回树的大小 即存储的字符串数量
func (n *Node) Size() int {
	r := 0
	for _, c := range n.children {
		r += c.Size() // 递归计算大小
	}
	if n.isWord {
		r++ // 如果是叶子节点，大小加一
	}
	return r
}

// remove 移除字符串
func (n *Node) remove(s string) {
	if len(s) == 0 {
		return // 如果字符串为空，返回
	}
	next, ok := n, false
	for _, c := range s {
		next, ok = next.children[c]
		if !ok {
			return // 字符串不存在，返回
		}
	}
	next.isWord = false // 标记为非叶子节点
}

// Remove 批量移除字符串
func (n *Node) Remove(s ...string) {
	for _, ss := range s {
		n.remove(ss) // 调用 remove 方法
	}
}

// Compact 压缩树
func (n *Node) Compact() (remove bool) {
	for r, c := range n.children {
		if c.Compact() {
			delete(n.children, r) // 删除空的孩子节点
		}
	}
	return !n.isWord && len(n.children) == 0 // 返回是否可以删除当前节点
}
