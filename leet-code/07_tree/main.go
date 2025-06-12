package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 226 翻转二叉树: https://leetcode.cn/problems/invert-binary-tree/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

// 98. 验证二叉搜索树  https://leetcode.cn/problems/validate-binary-search-tree/description/
func isValidBST(root *TreeNode) bool {
	return isValid(root, math.Inf(-1), math.Inf(1))
}

func isValid(root *TreeNode, min, max float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)
	return v > min && v < max && isValid(root.Left, min, v) && isValid(root.Right, v, max)
}

func isValidBST1(root *TreeNode) bool {
	arr := []int{}
	inOrder(root, &arr)
	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
}

// 104. 二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 111. 二叉树的最小深度 https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//为什么要特别处理单边为空的情况？
	// 因为最小深度定义是到"叶子节点"的最短路径
	// 叶子节点定义：没有任何子节点的节点
	// 如果一个节点只有左子树或只有右子树，我们必须走到那个非空的子树去找叶子节点
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 589. N 叉树的前序遍历 https://leetcode.cn/problems/n-ary-tree-preorder-traversal/
type Node struct {
	Val      int
	Children []*Node
}

func dfs(root *Node, list *[]int) {
	if root != nil {
		*list = append(*list, root.Val)
		for _, child := range root.Children {
			dfs(child, list)
		}
	}
}
func preorder(root *Node) []int {
	list := []int{}
	dfs(root, &list)
	return list
}

// 102. 二叉树的层序遍历 https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

// 429 层序遍历 https://leetcode.cn/problems/n-ary-tree-level-order-traversal/description/

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	result := [][]int{}
	queue := []*Node{root}

	for len(queue) > 0 {
		size := len(queue)         // 当前层的节点数
		level := make([]int, size) // 当前层的节点值
		for i := 0; i < size; i++ {
			node := queue[0]  // 取出队列头部的节点
			queue = queue[1:] // 移除队列头部的节点
			level[i] = node.Val
			queue = append(queue, node.Children...)
		}
		result = append(result, level)
	}

	return result
}

// 105. 从前序与中序遍历序列构造二叉树 https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
func build(preorder []int, inorder []int, l1, r1, l2, r2 int) *TreeNode {
	if l1 > r1 {
		return nil
	}
	root := &TreeNode{Val: preorder[l1]}
	mid := l2
	for inorder[mid] != root.Val {
		mid++
	}
	root.Left = build(preorder, inorder, l1+1, l1+(mid-l2), l2, mid-1)
	root.Right = build(preorder, inorder, l1+(mid-l2)+1, r1, mid+1, r2)
	return root

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func buildTree1(preorder []int, inorder []int) *TreeNode {
	mp := make(map[int]int)
	for i := range inorder {
		mp[inorder[i]] = i
	}

	var build func(preLeft, inLeft, inRight int) *TreeNode
	build = func(preLeft, inLeft, inRight int) *TreeNode {
		if inLeft > inRight {
			return nil
		}

		val := preorder[preLeft]
		node := &TreeNode{Val: val}
		inMid := mp[val]
		length := inMid - inLeft + 1
		node.Left = build(preLeft+1, inLeft, inMid-1)
		node.Right = build(preLeft+length, inMid+1, inRight)
		return node
	}

	return build(0, 0, len(inorder)-1)
}

// 297. 二叉树的序列化与反序列化 https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/description/

type Codec struct {
	Builder strings.Builder
	Input   []string
}

func Constructor() Codec {
	return Codec{}
}
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		this.Builder.WriteString("#,")
		return ""
	}
	this.Builder.WriteString(strconv.Itoa(root.Val) + ",")
	this.serialize(root.Left)
	this.serialize(root.Right)
	return this.Builder.String()
}

func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	this.Input = strings.Split(data, ",")
	return this.deserializeHelper()
}

func (this *Codec) deserializeHelper() *TreeNode {
	if this.Input[0] == "#" {
		this.Input = this.Input[1:]
		return nil
	}

	val, _ := strconv.Atoi(this.Input[0])
	this.Input = this.Input[1:]
	return &TreeNode{
		Val:   val,
		Left:  this.deserializeHelper(),
		Right: this.deserializeHelper(),
	}
}

// 1245 树的直径 https://leetcode.cn/problems/tree-diameter/description/
// 树的直径是树中任意两个节点之间的最长路径的长度
// 543. 二叉树的直径 https://leetcode.cn/problems/diameter-of-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked

var diameter int

func diameterOfBinaryTree(root *TreeNode) int {
	diameter = 0
	depth(root) // 计算深度并更新直径
	return diameter
}

// depth 计算树的深度并更新直径
func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// 递归计算左子树和右子树的深度
	leftDepth := depth(node.Left)
	rightDepth := depth(node.Right)

	// 更新直径
	diameter = max(diameter, leftDepth+rightDepth)

	// 返回当前节点的深度
	return max(leftDepth, rightDepth) + 1
}

// max 返回两个整数中的较大者
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 236 最近公共祖先 https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	return root
}

// 437 路径总和 III https://leetcode.cn/problems/path-sum-iii/description/
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// 计算以当前节点为起点的路径和
	count := countPaths(root, targetSum)

	// 递归计算左子树和右子树的路径和
	count += pathSum(root.Left, targetSum)
	count += pathSum(root.Right, targetSum)

	return count
}

func countPaths(node *TreeNode, targetSum int) int {
	if node == nil {
		return 0
	}

	// 检查当前节点是否等于目标值
	count := 0
	if node.Val == targetSum {
		count = 1
	}

	// 递归计算左子树和右子树的路径和
	count += countPaths(node.Left, targetSum-node.Val)
	count += countPaths(node.Right, targetSum-node.Val)

	return count
}

func pathSum1(root *TreeNode, targetSum int) (ans int) {
	// 存储前缀和的出现次数，初始化时前缀和 0 出现 1 次
	preSum := map[int64]int{0: 1}

	// 定义 DFS 函数，node 是当前节点，curr 是从根到当前节点的路径和
	var dfs func(*TreeNode, int64)
	dfs = func(node *TreeNode, curr int64) {
		if node == nil {
			return
		}

		// 将当前节点的值加到路径和中
		curr += int64(node.Val)

		// 如果存在一个前缀和，使得 curr - 该前缀和 = targetSum
		// 说明找到了一条符合要求的路径
		ans += preSum[curr-int64(targetSum)]

		// 将当前前缀和加入 map
		preSum[curr]++

		// 递归遍历左右子树
		dfs(node.Left, curr)
		dfs(node.Right, curr)

		// 回溯：将当前前缀和从 map 中删除
		preSum[curr]--
	}

	dfs(root, 0)
	return
}

// 124 二叉树中的最大路径和 https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 计算左子树和右子树的最大路径和
		left := max(0, dfs(node.Left))   // 如果左子树的路径和为负，则不考虑它
		right := max(0, dfs(node.Right)) // 同理

		// 更新全局最大路径和
		maxSum = max(maxSum, node.Val+left+right)

		// 返回当前节点的最大路径和
		return node.Val + max(left, right)
	}

	dfs(root)
	return maxSum
}

// 101. 对称二叉树 https://leetcode.cn/problems/symmetric-tree/description/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

// 114 二叉树展开为链表 https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 先递归处理左子树和右子树
	flatten(root.Left)
	flatten(root.Right)

	// 将左子树接到右子树上
	if root.Left != nil {
		right := root.Right // 保存右子树
		root.Right = root.Left
		root.Left = nil // 清空左子树

		// 找到新的右子树的最右节点，将原来的右子树接到它的右边
		current := root.Right
		for current.Right != nil {
			current = current.Right
		}
		current.Right = right
	}
}

func main() {

	// 构建一个示例树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	// 计算树的直径
	result := diameterOfBinaryTree(root)
	fmt.Println("树的直径是:", result) // 输出树的直径

	lowestCommonAncestor(root, root.Left, root.Right)
	pathSum(root, 8)
}
