package main

import (
	"fmt"
	"math"
	"math/bits"
)

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
func solveNQueens(n int) [][]string {
	var result [][]string       // 定义结果集
	cols := make(map[int]bool)  // 记录列
	diag1 := make(map[int]bool) // 主对角线 (\)
	diag2 := make(map[int]bool) // 副对角线 (/)
	// 定义当前棋盘
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	// 从第 0 行开始尝试放置皇后
	var backtrack func(row int)
	backtrack = func(row int) {
		// 如果行数等于 N，说明所有皇后已经成功放置，保存当前解
		if row == n {
			var solution []string
			for i := 0; i < n; i++ {
				solution = append(solution, string(board[i]))
			}
			result = append(result, solution)
			return
		}

		// 遍历当前行的每一列，尝试放置皇后
		for col := 0; col < n; col++ {
			// 检查该位置是否冲突
			if cols[col] || diag1[row-col] || diag2[row+col] {
				continue
			}

			// 放置皇后
			board[row][col] = 'Q'
			// 标记该位置已被占用
			cols[col] = true
			diag1[row-col] = true
			diag2[row+col] = true

			// 递归放置下一个皇后
			backtrack(row + 1)

			// 回溯，撤销当前选择
			board[row][col] = '.'
			cols[col] = false
			diag1[row-col] = false
			diag2[row+col] = false
		}
	}

	// 从第 0 行开始回溯
	backtrack(0)

	return result
}

// TODO 学完 位运算 再来
func solveNQueens1(n int) [][]string {
	res := make([][]string, 0)
	// 创建结果字符串的模板
	queens := make([]int, n)
	// 使用位运算记录占用情况
	columns := 0    // 列占用
	diagonals1 := 0 // 主对角线占用 (\)
	diagonals2 := 0 // 副对角线占用 (/)

	// 生成答案板的辅助函数
	generateBoard := func(queens []int, n int) []string {
		board := make([]string, n)
		row := make([]byte, n)
		for i := range row {
			row[i] = '.'
		}
		for i := 0; i < n; i++ {
			newRow := make([]byte, n)
			copy(newRow, row)
			newRow[queens[i]] = 'Q'
			board[i] = string(newRow)
		}
		return board
	}

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			board := generateBoard(queens, n)
			res = append(res, board)
			return
		}

		availablePositions := ((1 << n) - 1) &
			(^(columns | (diagonals1 >> row) | (diagonals2 >> (n - 1 - row))))

		for availablePositions != 0 {
			position := availablePositions & (-availablePositions)             // 获取最低位的1
			availablePositions = availablePositions & (availablePositions - 1) // 清除最低位的1
			column := bits.TrailingZeros(uint(position))

			queens[row] = column
			columns |= position
			diagonals1 |= position << row
			diagonals2 |= position << (n - 1 - row)

			backtrack(row + 1)

			columns &^= position
			diagonals1 &^= position << row
			diagonals2 &^= position << (n - 1 - row)
		}
	}

	backtrack(0)
	return res
}

// 200 岛屿数量 ：https://leetcode.cn/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	// 如果网格为空，直接返回 0
	if len(grid) == 0 {
		return 0
	}
	count := 0 // 用来统计岛屿数量

	// 遍历整个网格
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 如果当前格子是陆地（'1'），则认为发现了一个新的岛屿
			if grid[i][j] == '1' {
				count++         // 岛屿数量增加
				dfs(grid, i, j) // 对当前岛屿的所有陆地部分进行深度优先搜索，标记访问过的陆地为水（'0'）
			}
		}
	}
	return count // 返回岛屿的数量
}

// dfs 深度优先搜索，用于遍历并标记与当前陆地连接的所有陆地部分
func dfs(grid [][]byte, i, j int) {
	// 如果当前坐标超出边界或者已经是水（'0'），则返回
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}

	// 将当前的陆地（'1'）标记为水（'0'），表示已访问
	grid[i][j] = '0'

	// 递归地搜索上下左右四个方向的邻接格子
	dfs(grid, i-1, j) // 向上搜索
	dfs(grid, i+1, j) // 向下搜索
	dfs(grid, i, j-1) // 向左搜索
	dfs(grid, i, j+1) // 向右搜索
}

// 130 被围绕的区域 ： https://leetcode.cn/problems/surrounded-regions/
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		// 检查边界条件和是否是 'O'
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != 'O' {
			return
		}
		// 标记为安全区域
		board[r][c] = 'S' // S 表示安全的 O
		// 递归四个方向
		dfs(r+1, c) // 下
		dfs(r-1, c) // 上
		dfs(r, c+1) // 右
		dfs(r, c-1) // 左
	}

	// 从边界搜索所有的 O
	// 上边界和下边界
	for c := 0; c < cols; c++ {
		if board[0][c] == 'O' {
			dfs(0, c)
		}
		if board[rows-1][c] == 'O' {
			dfs(rows-1, c)
		}
	}

	// 左边界和右边界
	for r := 0; r < rows; r++ {
		if board[r][0] == 'O' {
			dfs(r, 0)
		}
		if board[r][cols-1] == 'O' {
			dfs(r, cols-1)
		}
	}

	// 处理剩下的 O 和 S，转换为 X 和 O
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X' // 被围绕的 O 转为 X
			}
			if board[r][c] == 'S' {
				board[r][c] = 'O' // 安全的 O 还原
			}
		}
	}
}

// 329 最长递增序列：https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/

// longestIncreasingPath 函数计算给定矩阵中最长递增路径的长度
func longestIncreasingPath(matrix [][]int) int {
	// 创建缓存数组，用于存储每个位置的最长递增路径长度
	cache, res := make([][]int, len(matrix)), 0
	for i := 0; i < len(cache); i++ {
		cache[i] = make([]int, len(matrix[0])) // 初始化缓存数组
	}

	// 遍历矩阵中的每个元素
	for i, v := range matrix {
		for j := range v {
			// 从当前元素开始搜索最长递增路径
			searchPath(matrix, cache, math.MinInt64, i, j)
			// 更新结果，取当前最长路径和之前结果的最大值
			res = max(res, cache[i][j])
		}
	}
	return res // 返回最长递增路径的长度
}

// 定义四个方向的移动：上、右、下、左
var dir = [][]int{
	{-1, 0}, // 上
	{0, 1},  // 右
	{1, 0},  // 下
	{0, -1}, // 左
}

// searchPath 函数递归地搜索从 (x, y) 开始的最长递增路径
func searchPath(board, cache [][]int, lastNum, x, y int) int {
	// 如果当前元素小于等于上一个元素，返回 0（不构成递增路径）
	if board[x][y] <= lastNum {
		return 0
	}
	// 如果当前元素的最长路径已经计算过，直接返回缓存值
	if cache[x][y] > 0 {
		return cache[x][y]
	}

	count := 1 // 初始化当前路径长度为 1（当前元素本身）
	// 遍历四个方向
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0] // 计算新位置的 x 坐标
		ny := y + dir[i][1] // 计算新位置的 y 坐标
		// 检查新位置是否在矩阵内
		if isInIntBoard(board, nx, ny) {
			// 递归搜索新位置的最长递增路径，并更新当前路径长度
			count = max(count, searchPath(board, cache, board[x][y], nx, ny)+1)
		}
	}
	cache[x][y] = count // 将当前元素的最长路径长度存入缓存
	return count        // 返回当前元素的最长递增路径长度
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// isInIntBoard 函数检查给定坐标 (x, y) 是否在矩阵的有效范围内
func isInIntBoard(board [][]int, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func main() {
	fmt.Println(letterCombinations("234"))
	hasPathSum(&TreeNode{}, 0)
	pathSum(&TreeNode{}, 0)
	kthSmallest(&TreeNode{}, 0)
	numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}})
	solveNQueens(4)
	solveNQueens1(4)

	solve([][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	})

	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	result1 := longestIncreasingPath(matrix)
	fmt.Println(result1) // 输出 4
}
