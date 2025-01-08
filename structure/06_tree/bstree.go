package tree

import "learn-go/structure/constraints"

var _ Node[int] = &BSNode[int]{}

// BSNode 是二叉搜索树的节点结构
type BSNode[T constraints.Ordered] struct {
	key    T          // 节点存储的值
	parent *BSNode[T] // 父节点指针
	left   *BSNode[T] // 左子节点指针
	right  *BSNode[T] // 右子节点指针
}

func (n *BSNode[T]) Key() T {
	return n.key
}

func (n *BSNode[T]) Parent() Node[T] {
	return n.parent
}

func (n *BSNode[T]) Left() Node[T] {
	return n.left
}

func (n *BSNode[T]) Right() Node[T] {
	return n.right
}

// BinarySearch 是二叉搜索树的主结构
type BinarySearch[T constraints.Ordered] struct {
	Root *BSNode[T] // 树的根节点
	_NIL *BSNode[T] // 哨兵节点，用于表示空节点
}

// NewBinarySearch 创建一个新的二叉搜索树
func NewBinarySearch[T constraints.Ordered]() *BinarySearch[T] {
	return &BinarySearch[T]{
		Root: nil,
		_NIL: nil,
	}
}

func (t *BinarySearch[T]) Empty() bool {
	return t.Root == t._NIL
}

// Push 向树中插入一个或多个值
func (t *BinarySearch[T]) Push(keys ...T) {
	for _, key := range keys {
		t.pushHelper(t.Root, key)
	}
}

// Delete 从树中删除一个值
func (t *BinarySearch[T]) Delete(val T) bool {
	node, ok := t.Get(val)
	if !ok {
		return false
	}
	t.deleteHelper(node.(*BSNode[T]))
	return true
}

func (t *BinarySearch[T]) Get(key T) (Node[T], bool) {
	return searchTreeHelper[T](t.Root, t._NIL, key)
}

func (t *BinarySearch[T]) Has(key T) bool {
	_, ok := searchTreeHelper[T](t.Root, t._NIL, key)
	return ok
}

func (t *BinarySearch[T]) PreOrder() []T {
	traversal := make([]T, 0)
	preOrderRecursive[T](t.Root, t._NIL, &traversal)
	return traversal
}

func (t *BinarySearch[T]) InOrder() []T {
	return inOrderHelper[T](t.Root, t._NIL)
}

func (t *BinarySearch[T]) PostOrder() []T {
	traversal := make([]T, 0)
	postOrderRecursive[T](t.Root, t._NIL, &traversal)
	return traversal
}

func (t *BinarySearch[T]) LevelOrder() []T {
	traversal := make([]T, 0)
	levelOrderHelper[T](t.Root, t._NIL, &traversal)
	return traversal
}

func (t *BinarySearch[T]) AccessNodesByLayer() [][]T {
	return accessNodeByLayerHelper[T](t.Root, t._NIL)
}

func (t *BinarySearch[T]) Depth() int {
	return calculateDepth[T](t.Root, t._NIL, 0)
}

func (t *BinarySearch[T]) Max() (T, bool) {
	ret := maximum[T](t.Root, t._NIL)
	if ret == t._NIL {
		var dft T
		return dft, false
	}
	return ret.Key(), true
}

func (t *BinarySearch[T]) Min() (T, bool) {
	ret := minimum[T](t.Root, t._NIL)
	if ret == t._NIL {
		var dft T
		return dft, false
	}
	return ret.Key(), true
}

// Predecessor 用来查找二叉搜索树中某个节点的前驱节点（predecessor）的值
func (t *BinarySearch[T]) Predecessor(key T) (T, bool) {
	node, ok := searchTreeHelper[T](t.Root, t._NIL, key)
	if !ok {
		var dft T
		return dft, ok
	}
	return predecessorHelper[T](node, t._NIL)
}

func (t *BinarySearch[T]) Successor(key T) (T, bool) {
	node, ok := searchTreeHelper[T](t.Root, t._NIL, key)
	if !ok {
		var dft T
		return dft, ok
	}
	return successorHelper[T](node, t._NIL)
}

// pushHelper 是插入操作的辅助函数
func (t *BinarySearch[T]) pushHelper(x *BSNode[T], val T) {
	y := t._NIL // 用于跟踪父节点

	// 查找插入位置
	for x != t._NIL {
		y = x
		switch {
		case val < x.Key():
			x = x.left
		case val > x.Key():
			x = x.right
		default: // 如果值已存在，直接返回
			return
		}
	}

	// 创建新节点
	z := &BSNode[T]{
		key:    val,
		left:   t._NIL,
		right:  t._NIL,
		parent: y,
	}

	// 将新节点连接到树中
	if y == t._NIL {
		t.Root = z // 树为空时，设置根节点
	} else if val < y.key {
		y.left = z // 作为左子节点
	} else {
		y.right = z // 作为右子节点
	}
}

// deleteHelper 是删除操作的辅助函数
func (t *BinarySearch[T]) deleteHelper(z *BSNode[T]) {
	switch {
	case z.left == t._NIL: // 没有左子树
		t.transplant(z, z.right)
	case z.right == t._NIL: // 没有右子树
		t.transplant(z, z.left)
	default: // 有两个子节点
		// 找到右子树中的最小节点作为后继
		y := minimum[T](z.right, t._NIL).(*BSNode[T])
		if y.parent != z {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}

		// 用后继节点替换被删除的节点
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
	}
}

// transplant 用于在删除操作中替换子树
func (t *BinarySearch[T]) transplant(u, v *BSNode[T]) {
	switch {
	case u.parent == t._NIL: // u是根节点
		t.Root = v
	case u == u.parent.left: // u是左子节点
		u.parent.left = v
	default: // u是右子节点
		u.parent.right = v
	}

	// 更新父节点引用
	if v != t._NIL {
		v.parent = u.parent
	}
}
