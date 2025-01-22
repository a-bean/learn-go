package main

import "fmt"

// 17 电话号码的字母组合 https://leetcode.cn/problems/letter-combinations-of-a-phone-number/

var (
	letterMap = []string{
		"",     //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	res = []string{}
)

// 解法一 DFS
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res = []string{}
	findCombination(&digits, 0, "")
	return res
}

func findCombination(digits *string, index int, s string) {
	if index == len(*digits) {
		res = append(res, s)
		return
	}
	num := (*digits)[index]
	letter := letterMap[num-'0']
	for i := 0; i < len(letter); i++ {
		findCombination(digits, index+1, s+string(letter[i]))
	}
}

// 112 路径总和 : https://leetcode.cn/problems/path-sum/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

// 113 路径总和 II: https://leetcode.cn/problems/path-sum-ii/
func pathSum(root *TreeNode, sum int) [][]int {
	var slice [][]int
	slice = findPath(root, sum, slice, []int(nil))
	return slice
}

func findPath(n *TreeNode, sum int, slice [][]int, stack []int) [][]int {
	if n == nil {
		return slice
	}
	sum -= n.Val
	stack = append(stack, n.Val)
	if sum == 0 && n.Left == nil && n.Right == nil {
		slice = append(slice, append([]int{}, stack...))
		stack = stack[:len(stack)-1]
	}
	slice = findPath(n.Left, sum, slice, stack)
	slice = findPath(n.Right, sum, slice, stack)
	return slice
}

// 230  二叉搜索树中第 K 小的元素: https://leetcode.cn/problems/kth-smallest-element-in-a-bst/
func kthSmallest(root *TreeNode, k int) int {
	ans, count := 0, 0
	inOrder(root, k, &count, &ans)
	return ans
}

func inOrder(root *TreeNode, k int, count *int, ans *int) {

	if root != nil {
		inOrder(root.Left, k, count, ans)
		*count++
		if k == *count {
			*ans = root.Val
			return
		}
		inOrder(root.Right, k, count, ans)
	}
}

// 235 二叉搜索树的最近公共祖先: https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/description
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p == nil || q == nil || root == nil {
		return nil
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

// 51 n皇后 ： https://leetcode.cn/problems/n-queens/
// TODO
func main() {
	fmt.Println(letterCombinations("234"))
	hasPathSum(&TreeNode{}, 0)
	pathSum(&TreeNode{}, 0)
}
