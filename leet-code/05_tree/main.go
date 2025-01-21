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

// 226 翻转二叉树 https://leetcode.cn/problems/invert-binary-tree/description/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	invertTree(root.Left)
	invertTree(root.Right)
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
	return v > min && v < max && isValid(root.Left, min, v) && isValid(root.Left, v, max)
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

//429 层序遍历 https://leetcode.cn/problems/n-ary-tree-level-order-traversal/description/

func levelOrder(root *Node) [][]int {
	var res [][]int
	var temp []int
	if root == nil {
		return res
	}
	queue := []*Node{root, nil}
	for len(queue) > 1 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			queue = append(queue, nil)
			res = append(res, temp)
			temp = []int{}
		} else {
			temp = append(temp, node.Val)
			if len(node.Children) > 0 {
				queue = append(queue, node.Children...)
			}
		}
	}
	res = append(res, temp)
	return res
}

// 105. 从前序与中序遍历序列构造二叉树 https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
func build(preorder []int, inorder []int, l1, r1, l2, r2 int) *TreeNode {
	if l1 > r2 {
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
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
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
}
