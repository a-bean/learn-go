package tree

import (
	"learn-go/structure/constraints"
)

// Node 定义了树节点的通用接口，支持泛型类型T（必须满足Ordered约束）
type Node[T constraints.Ordered] interface {
	Key() T
	Parent() Node[T]
	Left() Node[T]
	Right() Node[T]
}

// accessNodeByLayerHelper 实现树的层序遍历，返回每层节点值的二维切片
func accessNodeByLayerHelper[T constraints.Ordered](root, nilNode Node[T]) [][]T {
	if root == nilNode {
		return [][]T{}
	}
	var q []Node[T]
	var n Node[T]
	var idx = 0
	q = append(q, root)
	var res [][]T

	for len(q) != 0 {
		res = append(res, []T{})
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			n, q = q[0], q[1:]
			res[idx] = append(res[idx], n.Key())
			if n.Left() != nilNode {
				q = append(q, n.Left())
			}
			if n.Right() != nilNode {
				q = append(q, n.Right())
			}
		}
		idx++
	}
	return res
}

// searchTreeHelper 在二叉搜索树中查找指定值，返回找到的节点和是否存在的标志
func searchTreeHelper[T constraints.Ordered](node, nilNode Node[T], key T) (Node[T], bool) {
	if node == nilNode {
		return node, false
	}

	if key == node.Key() {
		return node, true
	}
	if key < node.Key() {
		return searchTreeHelper(node.Left(), nilNode, key)
	}
	return searchTreeHelper(node.Right(), nilNode, key)
}

// inOrderHelper 实现树的中序遍历（非递归方式）
// 使用栈来模拟递归过程，按照 左-根-右 的顺序访问节点 从小到大返回节点
func inOrderHelper[T constraints.Ordered](node, nilNode Node[T]) []T {
	var stack []Node[T]
	var ret []T

	for node != nilNode || len(stack) > 0 {
		// 将所有左子节点入栈
		for node != nilNode {
			stack = append(stack, node)
			node = node.Left()
		}
		// 处理栈顶节点
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Key())
		// 处理右子节点
		node = node.Right()
	}

	return ret
}

// preOrderRecursive 实现树的前序遍历（递归方式）
// 按照 根-左-右 的顺序访问节点
func preOrderRecursive[T constraints.Ordered](n, nilNode Node[T], traversal *[]T) {
	if n == nilNode {
		return
	}

	*traversal = append(*traversal, n.Key())
	preOrderRecursive(n.Left(), nilNode, traversal)
	preOrderRecursive(n.Right(), nilNode, traversal)

}

// postOrderRecursive 实现树的后序遍历（递归方式）
// 按照 左-右-根 的顺序访问节点
func postOrderRecursive[T constraints.Ordered](n, nilNode Node[T], traversal *[]T) {
	if n == nilNode {
		return
	}

	postOrderRecursive(n.Left(), nilNode, traversal)
	postOrderRecursive(n.Right(), nilNode, traversal)
	*traversal = append(*traversal, n.Key())
}

// calculateDepth 计算树的深度
// 递归计算左右子树的最大深度，返回较大值加1
func calculateDepth[T constraints.Ordered](n, nilNode Node[T], depth int) int {
	if n == nilNode {
		return depth
	}

	return maxInt(calculateDepth(n.Left(), nilNode, depth+1), calculateDepth(n.Right(), nilNode, depth+1))
}

// minimum 查找树中的最小值节点
// 在二叉搜索树中，最左边的叶子节点即为最小值
func minimum[T constraints.Ordered](node, nilNode Node[T]) Node[T] {
	if node == nilNode {
		return node
	}

	for node.Left() != nilNode {
		node = node.Left()
	}
	return node
}

// maximum 查找树中的最大值节点
// 在二叉搜索树中，最右边的叶子节点即为最大值
func maximum[T constraints.Ordered](node, nilNode Node[T]) Node[T] {
	if node == nilNode {
		return node
	}

	for node.Right() != nilNode {
		node = node.Right()
	}
	return node
}

// levelOrderHelper 实现树的层序遍历
// 使用队列来按层访问节点
func levelOrderHelper[T constraints.Ordered](root, nilNode Node[T], traversal *[]T) {
	var q []Node[T] // queue
	var tmp Node[T]

	q = append(q, root)

	for len(q) != 0 {
		tmp, q = q[0], q[1:]
		*traversal = append(*traversal, tmp.Key())
		if tmp.Left() != nilNode {
			q = append(q, tmp.Left())
		}

		if tmp.Right() != nilNode {
			q = append(q, tmp.Right())
		}
	}
}

// predecessorHelper 查找指定节点的前驱节点
// 前驱节点是中序遍历中当前节点的前一个节点
func predecessorHelper[T constraints.Ordered](node, nilNode Node[T]) (T, bool) {
	// 如果有左子树，则前驱是左子树中的最大值
	if node.Left() != nilNode {
		return maximum(node.Left(), nilNode).Key(), true
	}
	// 如果没有左子树，则向上查找第一个将当前节点作为右子树的祖先节点
	p := node.Parent()
	for p != nilNode && node == p.Left() {
		node = p
		p = p.Parent()
	}

	if p == nilNode {
		var dft T
		return dft, false
	}
	return p.Key(), true
}

// successorHelper 查找指定节点的后继节点
// 后继节点是中序遍历中当前节点的后一个节点
func successorHelper[T constraints.Ordered](node, nilNode Node[T]) (T, bool) {
	// 如果有右子树，则后继是右子树中的最小值
	if node.Right() != nilNode {
		return minimum(node.Right(), nilNode).Key(), true
	}

	// 如果没有右子树，则向上查找第一个将当前节点作为左子树的祖先节点
	p := node.Parent()
	for p != nilNode && node == p.Right() {
		node = p
		p = p.Parent()
	}

	if p == nilNode {
		var dft T
		return dft, false
	}
	return p.Key(), true
}
