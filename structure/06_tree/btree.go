package tree

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"
)

// Item 接口定义了可以存储在B树中的元素类型
// 任何实现此接口的类型都必须提供Less方法来比较元素大小
type Item interface {
	Less(than Item) bool
}

const (
	// DefaultFreeListSize 定义了默认空闲列表的大小
	DefaultFreeListSize = 32
)

var (
	// nilItems 是一个预分配的空items切片，用于重置items
	nilItems = make(items, 16)
	// nilChildren 是一个预分配的空children切片，用于重置children
	nilChildren = make(children, 16)
)

// FreeList 是一个节点对象池，用于减少内存分配
type FreeList struct {
	mu       sync.Mutex // 互斥锁，保证并发安全
	freelist []*node    // 存储可重用的节点
}

// NewFreeList 创建一个指定大小的FreeList
func NewFreeList(size int) *FreeList {
	return &FreeList{freelist: make([]*node, 0, size)}
}

// newNode 从空闲列表中获取一个节点，如果列表为空则创建新节点
func (f *FreeList) newNode() (n *node) {
	f.mu.Lock()
	index := len(f.freelist) - 1
	if index < 0 {
		f.mu.Unlock()
		return new(node) // 空闲列表为空，创建新节点
	}
	n = f.freelist[index]
	f.freelist[index] = nil
	f.freelist = f.freelist[:index]
	f.mu.Unlock()
	return
}

// freeNode 将不再使用的节点放回空闲列表
func (f *FreeList) freeNode(n *node) (out bool) {
	f.mu.Lock()
	if len(f.freelist) < cap(f.freelist) {
		f.freelist = append(f.freelist, n)
		out = true
	}
	f.mu.Unlock()
	return
}

// ItemIterator 是一个函数类型，用于遍历B树中的元素
// 返回false可以停止遍历
type ItemIterator func(i Item) bool

// New 创建一个新的B树，使用默认大小的空闲列表
func New(degree int) *BTree {
	return NewWithFreeList(degree, NewFreeList(DefaultFreeListSize))
}

// NewWithFreeList 创建一个新的B树，使用指定的空闲列表
func NewWithFreeList(degree int, f *FreeList) *BTree {
	if degree <= 1 {
		panic("bad degree")
	}
	return &BTree{
		degree: degree,
		cow:    &copyOnWriteContext{freelist: f},
	}
}

// items 是Item的切片类型
type items []Item

// insertAt 在指定索引位置插入一个元素
func (s *items) insertAt(index int, item Item) {
	*s = append(*s, nil)
	if index < len(*s) {
		copy((*s)[index+1:], (*s)[index:])
	}
	(*s)[index] = item
}

// removeAt 移除指定索引位置的元素并返回它
func (s *items) removeAt(index int) Item {
	item := (*s)[index]
	copy((*s)[index:], (*s)[index+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return item
}

// pop 移除并返回最后一个元素
func (s *items) pop() (out Item) {
	index := len(*s) - 1
	out = (*s)[index]
	(*s)[index] = nil
	*s = (*s)[:index]
	return
}

// truncate 截断切片到指定长度，并清空被截断的部分
func (s *items) truncate(index int) {
	var toClear items
	*s, toClear = (*s)[:index], (*s)[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, nilItems):]
	}
}

// find 查找元素在切片中的位置
// 返回元素应该插入的位置和是否找到元素
func (s items) find(item Item) (index int, found bool) {
	i := sort.Search(len(s), func(i int) bool {
		return item.Less(s[i])
	})
	if i > 0 && !s[i-1].Less(item) {
		return i - 1, true
	}
	return i, false
}

// children 是node指针的切片类型
type children []*node

// insertAt 在指定索引位置插入一个子节点
func (s *children) insertAt(index int, n *node) {
	*s = append(*s, nil)
	if index < len(*s) {
		copy((*s)[index+1:], (*s)[index:])
	}
	(*s)[index] = n
}

// removeAt 移除指定索引位置的子节点并返回它
func (s *children) removeAt(index int) *node {
	n := (*s)[index]
	copy((*s)[index:], (*s)[index+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return n
}

// pop 移除并返回最后一个子节点
func (s *children) pop() (out *node) {
	index := len(*s) - 1
	out = (*s)[index]
	(*s)[index] = nil
	*s = (*s)[:index]
	return
}

// truncate 截断切片到指定长度，并清空被截断的部分
func (s *children) truncate(index int) {
	var toClear children
	*s, toClear = (*s)[:index], (*s)[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, nilChildren):]
	}
}

// node 是B树的节点结构
type node struct {
	items    items               // 存储在节点中的元素
	children children            // 子节点指针
	cow      *copyOnWriteContext // 写时复制上下文
}

// mutableFor 确保节点对指定的cow上下文是可变的
// 如果节点已经属于该上下文，则直接返回；否则创建一个副本
func (n *node) mutableFor(cow *copyOnWriteContext) *node {
	if n.cow == cow {
		return n
	}
	out := cow.newNode()
	if cap(out.items) >= len(n.items) {
		out.items = out.items[:len(n.items)]
	} else {
		out.items = make(items, len(n.items), cap(n.items))
	}
	copy(out.items, n.items)
	// 复制子节点
	if cap(out.children) >= len(n.children) {
		out.children = out.children[:len(n.children)]
	} else {
		out.children = make(children, len(n.children), cap(n.children))
	}
	copy(out.children, n.children)
	return out
}

// mutableChild 确保指定索引的子节点是可变的
func (n *node) mutableChild(i int) *node {
	c := n.children[i].mutableFor(n.cow)
	n.children[i] = c
	return c
}

// split 将节点在指定位置分裂成两个节点
// 返回中间元素和新创建的右侧节点
func (n *node) split(i int) (Item, *node) {
	item := n.items[i]
	next := n.cow.newNode()
	next.items = append(next.items, n.items[i+1:]...)
	n.items.truncate(i)
	if len(n.children) > 0 {
		next.children = append(next.children, n.children[i+1:]...)
		n.children.truncate(i + 1)
	}
	return item, next
}

// maybeSplitChild 检查并在必要时分裂子节点
func (n *node) maybeSplitChild(i, maxItems int) bool {
	if len(n.children[i].items) < maxItems {
		return false
	}
	first := n.mutableChild(i)
	item, second := first.split(maxItems / 2)
	n.items.insertAt(i, item)
	n.children.insertAt(i+1, second)
	return true
}

// insert 在节点中插入元素
// 如果元素已存在，则替换并返回原元素；否则返回nil
func (n *node) insert(item Item, maxItems int) Item {
	i, found := n.items.find(item)
	if found {
		out := n.items[i]
		n.items[i] = item
		return out
	}
	if len(n.children) == 0 {
		n.items.insertAt(i, item)
		return nil
	}
	if n.maybeSplitChild(i, maxItems) {
		inTree := n.items[i]
		switch {
		case item.Less(inTree):
			// 不变，我们需要第一个分裂节点
		case inTree.Less(item):
			i++ // 我们需要第二个分裂节点
		default:
			out := n.items[i]
			n.items[i] = item
			return out
		}
	}
	return n.mutableChild(i).insert(item, maxItems)
}

// get 在节点中查找元素
func (n *node) get(key Item) Item {
	i, found := n.items.find(key)
	if found {
		return n.items[i]
	} else if len(n.children) > 0 {
		return n.children[i].get(key)
	}
	return nil
}

// min 返回节点子树中的最小元素
func min(n *node) Item {
	if n == nil {
		return nil
	}
	for len(n.children) > 0 {
		n = n.children[0]
	}
	if len(n.items) == 0 {
		return nil
	}
	return n.items[0]
}

// max 返回节点子树中的最大元素
func max(n *node) Item {
	if n == nil {
		return nil
	}
	for len(n.children) > 0 {
		n = n.children[len(n.children)-1]
	}
	if len(n.items) == 0 {
		return nil
	}
	return n.items[len(n.items)-1]
}

// toRemove 定义了删除操作的类型
type toRemove int

const (
	removeItem toRemove = iota // 删除指定元素
	removeMin                  // 删除子树中的最小元素
	removeMax                  // 删除子树中的最大元素
)

// remove 从节点中删除元素
func (n *node) remove(item Item, minItems int, typ toRemove) Item {
	var i int
	var found bool
	switch typ {
	case removeMax:
		if len(n.children) == 0 {
			return n.items.pop()
		}
		i = len(n.items)
	case removeMin:
		if len(n.children) == 0 {
			return n.items.removeAt(0)
		}
		i = 0
	case removeItem:
		i, found = n.items.find(item)
		if len(n.children) == 0 {
			if found {
				return n.items.removeAt(i)
			}
			return nil
		}
	default:
		panic("invalid type")
	}
	if len(n.children[i].items) <= minItems {
		return n.growChildAndRemove(i, item, minItems, typ)
	}
	child := n.mutableChild(i)

	if found {
		// 如果在当前节点找到了元素，用右子树的最小元素替换它
		out := n.items[i]
		n.items[i] = child.remove(nil, minItems, removeMax)
		return out
	}

	return child.remove(item, minItems, typ)
}

// growChildAndRemove 在删除前确保子节点有足够的元素
func (n *node) growChildAndRemove(i int, item Item, minItems int, typ toRemove) Item {
	if i > 0 && len(n.children[i-1].items) > minItems {
		// 从左侧子节点借用元素
		child := n.mutableChild(i)
		stealFrom := n.mutableChild(i - 1)
		stolenItem := stealFrom.items.pop()
		child.items.insertAt(0, n.items[i-1])
		n.items[i-1] = stolenItem
		if len(stealFrom.children) > 0 {
			child.children.insertAt(0, stealFrom.children.pop())
		}
	} else if i < len(n.items) && len(n.children[i+1].items) > minItems {
		// 从右侧子节点借用元素
		child := n.mutableChild(i)
		stealFrom := n.mutableChild(i + 1)
		stolenItem := stealFrom.items.removeAt(0)
		child.items = append(child.items, n.items[i])
		n.items[i] = stolenItem
		if len(stealFrom.children) > 0 {
			child.children = append(child.children, stealFrom.children.removeAt(0))
		}
	} else {
		// 合并子节点
		if i >= len(n.items) {
			i--
		}
		child := n.mutableChild(i)
		// 与右侧子节点合并
		mergeItem := n.items.removeAt(i)
		mergeChild := n.children.removeAt(i + 1).mutableFor(n.cow)
		child.items = append(child.items, mergeItem)
		child.items = append(child.items, mergeChild.items...)
		child.children = append(child.children, mergeChild.children...)
		n.cow.freeNode(mergeChild)
	}
	return n.remove(item, minItems, typ)
}

// direction 定义了遍历的方向
type direction int

const (
	descend = direction(-1) // 降序
	// ascend 表示升序遍历
	ascend = direction(+1)
)

// iterate 遍历节点及其子节点中的元素
// dir: 遍历方向（升序或降序）
// start: 起始元素（可为nil）
// stop: 结束元素（可为nil）
// includeStart: 是否包含起始元素
// hit: 是否已经命中起始元素
// iter: 遍历回调函数
func (n *node) iterate(dir direction, start, stop Item, includeStart bool, hit bool, iter ItemIterator) (bool, bool) {
	var ok, found bool
	var index int
	switch dir {
	case ascend: // 升序遍历
		if start != nil {
			index, _ = n.items.find(start)
		}
		for i := index; i < len(n.items); i++ {
			if len(n.children) > 0 {
				if hit, ok = n.children[i].iterate(dir, start, stop, includeStart, hit, iter); !ok {
					return hit, false
				}
			}
			if !includeStart && !hit && start != nil && !start.Less(n.items[i]) {
				hit = true
				continue
			}
			hit = true
			if stop != nil && !n.items[i].Less(stop) {
				return hit, false
			}
			if !iter(n.items[i]) {
				return hit, false
			}
		}
		if len(n.children) > 0 {
			if hit, ok = n.children[len(n.children)-1].iterate(dir, start, stop, includeStart, hit, iter); !ok {
				return hit, false
			}
		}
	case descend: // 降序遍历
		if start != nil {
			index, found = n.items.find(start)
			if !found {
				index = index - 1
			}
		} else {
			index = len(n.items) - 1
		}
		for i := index; i >= 0; i-- {
			if start != nil && !n.items[i].Less(start) {
				if !includeStart || hit || start.Less(n.items[i]) {
					continue
				}
			}
			if len(n.children) > 0 {
				if hit, ok = n.children[i+1].iterate(dir, start, stop, includeStart, hit, iter); !ok {
					return hit, false
				}
			}
			if stop != nil && !stop.Less(n.items[i]) {
				return hit, false
			}
			hit = true
			if !iter(n.items[i]) {
				return hit, false
			}
		}
		if len(n.children) > 0 {
			if hit, ok = n.children[0].iterate(dir, start, stop, includeStart, hit, iter); !ok {
				return hit, false
			}
		}
	}
	return hit, true
}

// print 打印节点及其子节点的内容，用于调试
func (n *node) print(w io.Writer, level int) {
	fmt.Fprintf(w, "%sNODE:%v\n", strings.Repeat("  ", level), n.items)
	for _, c := range n.children {
		c.print(w, level+1)
	}
}

// BTree 是B树的主结构
type BTree struct {
	degree int                 // B树的度，决定了节点可以包含的最大元素数
	length int                 // 树中元素的总数
	root   *node               // 根节点
	cow    *copyOnWriteContext // 写时复制上下文
}

// copyOnWriteContext 是写时复制的上下文
// 用于支持树的并发操作和克隆
type copyOnWriteContext struct {
	freelist *FreeList // 空闲节点列表
}

// Clone 创建树的一个副本
// 使用写时复制技术，初始时两棵树共享相同的节点
func (t *BTree) Clone() (t2 *BTree) {
	// 创建新的cow上下文
	cow1, cow2 := *t.cow, *t.cow
	out := *t
	t.cow = &cow1
	out.cow = &cow2
	return &out
}

// maxItems 返回节点可以包含的最大元素数
func (t *BTree) maxItems() int {
	return t.degree*2 - 1
}

// minItems 返回非根节点必须包含的最小元素数
func (t *BTree) minItems() int {
	return t.degree - 1
}

// newNode 从空闲列表中获取一个节点并设置其cow上下文
func (c *copyOnWriteContext) newNode() (n *node) {
	n = c.freelist.newNode()
	n.cow = c
	return
}

// freeType 表示节点释放的结果类型
type freeType int

const (
	ftFreelistFull freeType = iota // 节点被释放（可被GC回收，未存储在空闲列表中）
	ftStored                       // 节点被存储在空闲列表中以供后续使用
	ftNotOwned                     // 节点被COW忽略，因为它属于另一个上下文
)

// freeNode 释放节点并返回释放结果
func (c *copyOnWriteContext) freeNode(n *node) freeType {
	if n.cow == c {
		// 清空节点以允许GC回收
		n.items.truncate(0)
		n.children.truncate(0)
		n.cow = nil
		if c.freelist.freeNode(n) {
			return ftStored
		} else {
			return ftFreelistFull
		}
	} else {
		return ftNotOwned
	}
}

// ReplaceOrInsert 插入元素到树中，如果元素已存在则替换并返回原元素
func (t *BTree) ReplaceOrInsert(item Item) Item {
	if item == nil {
		panic("nil item being added to BTree")
	}
	if t.root == nil {
		// 树为空，创建根节点
		t.root = t.cow.newNode()
		t.root.items = append(t.root.items, item)
		t.length++
		return nil
	} else {
		t.root = t.root.mutableFor(t.cow)
		if len(t.root.items) >= t.maxItems() {
			// 根节点已满，需要分裂
			item2, second := t.root.split(t.maxItems() / 2)
			oldroot := t.root
			t.root = t.cow.newNode()
			t.root.items = append(t.root.items, item2)
			t.root.children = append(t.root.children, oldroot, second)
		}
	}
	out := t.root.insert(item, t.maxItems())
	if out == nil {
		t.length++
	}
	return out
}

// Delete 从树中删除指定元素并返回它
func (t *BTree) Delete(item Item) Item {
	return t.deleteItem(item, removeItem)
}

// DeleteMin 删除并返回树中的最小元素
func (t *BTree) DeleteMin() Item {
	return t.deleteItem(nil, removeMin)
}

// DeleteMax 删除并返回树中的最大元素
func (t *BTree) DeleteMax() Item {
	return t.deleteItem(nil, removeMax)
}

// deleteItem 实现删除操作的通用方法
func (t *BTree) deleteItem(item Item, typ toRemove) Item {
	if t.root == nil || len(t.root.items) == 0 {
		return nil
	}
	t.root = t.root.mutableFor(t.cow)
	out := t.root.remove(item, t.minItems(), typ)
	if len(t.root.items) == 0 && len(t.root.children) > 0 {
		// 根节点为空但有子节点，将第一个子节点作为新的根
		oldroot := t.root
		t.root = t.root.children[0]
		t.cow.freeNode(oldroot)
	}
	if out != nil {
		t.length--
	}
	return out
}

// AscendRange 升序遍历指定范围内的元素
// 范围是 [greaterOrEqual, lessThan)
func (t *BTree) AscendRange(greaterOrEqual, lessThan Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(ascend, greaterOrEqual, lessThan, true, false, iterator)
}

// AscendLessThan 升序遍历小于pivot的所有元素
func (t *BTree) AscendLessThan(pivot Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(ascend, nil, pivot, false, false, iterator)
}

// AscendGreaterOrEqual 升序遍历大于等于pivot的所有元素
func (t *BTree) AscendGreaterOrEqual(pivot Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(ascend, pivot, nil, true, false, iterator)
}

// Ascend 升序遍历树中的所有元素
func (t *BTree) Ascend(iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(ascend, nil, nil, false, false, iterator)
}

// DescendRange 降序遍历指定范围内的元素
// 范围是 (greaterThan, lessOrEqual]
func (t *BTree) DescendRange(lessOrEqual, greaterThan Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(descend, lessOrEqual, greaterThan, true, false, iterator)
}

// DescendLessOrEqual 降序遍历小于等于pivot的所有元素
func (t *BTree) DescendLessOrEqual(pivot Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(descend, pivot, nil, true, false, iterator)
}

// DescendGreaterThan 降序遍历大于pivot的所有元素
func (t *BTree) DescendGreaterThan(pivot Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(descend, nil, pivot, false, false, iterator)
}

// Descend 降序遍历树中的所有元素
func (t *BTree) Descend(iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(descend, nil, nil, false, false, iterator)
}

// Get 获取与key匹配的元素
func (t *BTree) Get(key Item) Item {
	if t.root == nil {
		return nil
	}
	return t.root.get(key)
}

// Min 返回树中的最小元素
func (t *BTree) Min() Item {
	return min(t.root)
}

// Max 返回树中的最大元素
func (t *BTree) Max() Item {
	return max(t.root)
}

// Has 检查树中是否存在指定的元素
func (t *BTree) Has(key Item) bool {
	return t.Get(key) != nil
}

// Len 返回树中元素的数量
func (t *BTree) Len() int {
	return t.length
}

// Clear 清空树
// 如果addNodesToFreelist为true，则将节点添加到空闲列表中
func (t *BTree) Clear(addNodesToFreelist bool) {
	if t.root != nil && addNodesToFreelist {
		t.root.reset(t.cow)
	}
	t.root, t.length = nil, 0
}

// reset 递归重置节点及其子节点
func (n *node) reset(c *copyOnWriteContext) bool {
	for _, child := range n.children {
		if !child.reset(c) {
			return false
		}
	}
	return c.freeNode(n) != ftFreelistFull
}

// Int 是一个实现了Item接口的整数类型
type Int int

// Less 比较两个Int的大小
func (a Int) Less(b Item) bool {
	return a < b.(Int)
}
