package main

import (
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/invert-binary-tree/description/
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

// https://leetcode.cn/problems/validate-binary-search-tree/description/
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

// https://leetcode.cn/problems/n-ary-tree-preorder-traversal/
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

// 层序遍历 https://leetcode.cn/problems/n-ary-tree-level-order-traversal/description/

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

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
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

// https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/description/

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

// https://leetcode.cn/problems/tree-diameter/description/

func main() {
}
