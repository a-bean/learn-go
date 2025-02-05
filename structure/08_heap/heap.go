package heap

import (
	"errors"
	"learn-go/structure/constraints"
)

// Heap 结构体定义
// T 是堆中元素的类型
// heaps 是存储堆元素的切片
// lessFunc 是用于比较两个元素的函数，决定堆的顺序
type Heap[T any] struct {
	heaps    []T
	lessFunc func(a, b T) bool
}

// New 构造函数：用于内置可比较类型
// 返回一个最小堆，元素类型为 T
func New[T constraints.Ordered]() *Heap[T] {
	// 定义默认的比较函数：a < b
	less := func(a, b T) bool {
		return a < b
	}
	// 调用 NewAny 创建堆
	h, _ := NewAny[T](less)
	return h
}

// NewAny 构造函数：用于任意类型和自定义比较函数
// less 是用于比较两个元素的函数
// 返回一个堆实例和可能的错误（如果 less 为 nil）
func NewAny[T any](less func(a, b T) bool) (*Heap[T], error) {
	// 检查 less 函数是否为空
	if less == nil {
		return nil, errors.New("less func is necessary")
	}
	// 返回初始化后的堆实例
	return &Heap[T]{
		lessFunc: less,
	}, nil
}

// Push 方法：向堆中插入一个元素
// t 是要插入的元素
func (h *Heap[T]) Push(t T) {
	// 将元素添加到切片末尾
	h.heaps = append(h.heaps, t)
	// 上浮新元素以维护堆性质
	h.up(len(h.heaps) - 1)
}

// Top 方法：获取堆顶元素
// 返回堆顶元素（不删除）
func (h *Heap[T]) Top() T {
	return h.heaps[0]
}

// Pop 方法：删除堆顶元素
func (h *Heap[T]) Pop() {
	// 如果堆中只有一个元素，直接清空
	if len(h.heaps) <= 1 {
		h.heaps = nil
		return
	}
	// 将堆顶元素与最后一个元素交换
	h.swap(0, len(h.heaps)-1)
	// 删除最后一个元素（原堆顶）
	h.heaps = h.heaps[:len(h.heaps)-1]
	// 下沉新的堆顶元素以维护堆性质
	h.down(0)
}

// Empty 方法：判断堆是否为空
// 返回 true 如果堆为空，否则返回 false
func (h *Heap[T]) Empty() bool {
	return len(h.heaps) == 0
}

// Size 方法：获取堆中元素的数量
// 返回堆中元素的数量
func (h *Heap[T]) Size() int {
	return len(h.heaps)
}

// swap 方法：交换堆中两个位置的元素
// i, j 是要交换的元素索引
func (h *Heap[T]) swap(i, j int) {
	h.heaps[i], h.heaps[j] = h.heaps[j], h.heaps[i]
}

// up 方法：上浮操作，维护堆性质
// child 是要上浮的元素的索引
func (h *Heap[T]) up(child int) {
	// 如果 child 已经是根节点，停止上浮
	if child <= 0 {
		return
	}
	// 计算父节点索引
	parent := (child - 1) >> 1
	// 如果 child 不小于 parent，停止上浮
	if !h.lessFunc(h.heaps[child], h.heaps[parent]) {
		return
	}
	// 交换 child 和 parent
	h.swap(child, parent)
	// 递归上浮
	h.up(parent)
}

// down 方法：下沉操作，维护堆性质
// parent 是要下沉的元素的索引
func (h *Heap[T]) down(parent int) {
	// 初始化最小元素索引为 parent
	lessIdx := parent
	// 计算左右子节点索引
	lChild, rChild := (parent<<1)+1, (parent<<1)+2
	// 如果左子节点存在且小于当前最小元素，更新最小元素索引
	if lChild < len(h.heaps) && h.lessFunc(h.heaps[lChild], h.heaps[lessIdx]) {
		lessIdx = lChild
	}
	// 如果右子节点存在且小于当前最小元素，更新最小元素索引
	if rChild < len(h.heaps) && h.lessFunc(h.heaps[rChild], h.heaps[lessIdx]) {
		lessIdx = rChild
	}
	// 如果 parent 已经是最小元素，停止下沉
	if lessIdx == parent {
		return
	}
	// 交换 parent 和最小元素
	h.swap(lessIdx, parent)
	// 递归下沉
	h.down(lessIdx)
}
