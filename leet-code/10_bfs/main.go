package main

import (
	"fmt"
)

// 200 岛屿数量 ：https://leetcode.cn/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	// 遍历每一个单元格
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 如果找到一个陆地单元，则启动 BFS
			if grid[i][j] == '1' {
				count++
				bfs(grid, i, j)
			}
		}
	}
	return count
}

// bfs 实现广度优先搜索
func bfs(grid [][]byte, i, j int) {
	// 定义方向数组，表示上下左右的移动
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := [][2]int{{i, j}} // 使用队列来存储待处理的坐标

	// 将当前陆地单元标记为水
	grid[i][j] = '0'

	// 处理队列中的元素
	for len(queue) > 0 {
		// 取出队列中的第一个元素
		curr := queue[0]
		queue = queue[1:] // 删除第一项
		currI, currJ := curr[0], curr[1]

		// 遍历四个方向
		for _, d := range directions {
			newI, newJ := currI+d[0], currJ+d[1]
			if newI >= 0 && newJ >= 0 && newI < len(grid) && newJ < len(grid[0]) && grid[newI][newJ] == '1' {
				// 标记为水（访问过）
				grid[newI][newJ] = '0'
				queue = append(queue, [2]int{newI, newJ}) // 将新的陆地坐标加入队列
			}
		}
	}
}

// 433 最小基因变化：https://leetcode.cn/problems/minimum-genetic-mutation/
func minMutation(start string, end string, bank []string) int {
	// Step 1: 检查目标基因 (end) 是否存在于基因池中
	found := false
	bankSet := make(map[string]bool) // 使用 map 来存储基因池，方便 O(1) 查找

	// 遍历基因池，将每个基因加入 bankSet 中，并检查目标基因 (end) 是否存在
	for _, gene := range bank {
		bankSet[gene] = true
		if gene == end {
			found = true // 如果目标基因在基因池中，标记为 true
		}
	}

	// 如果目标基因不存在于基因池中，返回 -1，表示无法变换到目标基因
	if !found {
		return -1
	}

	// Step 2: 初始化 BFS
	queue := []string{start}         // 用队列来存储待处理的基因序列，初始化时将起始基因加入队列
	visited := make(map[string]bool) // 使用 map 来记录已经访问过的基因，避免重复处理
	visited[start] = true            // 标记起始基因为已访问
	steps := 0                       // 记录基因变化的步数

	// Step 3: 基因字符集（'A', 'C', 'G', 'T'）
	geneChars := []byte{'A', 'C', 'G', 'T'} // 基因变化时可选择的字符集

	// Step 4: BFS 处理基因变化
	for len(queue) > 0 { // 当队列中还有基因序列需要处理时，继续循环
		size := len(queue)          // 获取当前队列中的基因数量（即当前处理的层级）
		for i := 0; i < size; i++ { // 遍历当前层的所有基因
			current := queue[i] // 取出当前基因序列
			if current == end { // 如果当前基因等于目标基因，直接返回步骤数
				return steps
			}

			// Step 5: 尝试变换当前基因序列中的每个字符
			for j := 0; j < len(current); j++ { // 遍历基因序列中的每个字符
				for _, char := range geneChars { // 尝试用 geneChars 中的每个字符替换当前字符
					newGene := current[:j] + string(char) + current[j+1:] // 创建新的基因序列

					// Step 6: 检查新基因是否在基因池中，并且没有被访问过
					if _, ok := bankSet[newGene]; ok && !visited[newGene] {
						visited[newGene] = true        // 标记新基因为已访问
						queue = append(queue, newGene) // 将新基因加入队列，待下轮处理
					}
				}
			}
		}
		queue = queue[size:] // 移除当前层已经处理过的基因，处理下一层
		steps++              // 增加变换的步数
	}

	// 如果 BFS 遍历结束，仍未找到目标基因，返回 -1
	return -1
}

// 329 最长递增序列：https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/

// 329 矩阵中的最长递增路径：https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/
// 使用拓扑排序（Kahn算法）结合动态规划，时间复杂度O(mn)，空间复杂度O(mn)
func longestIncreasingPath(matrix [][]int) int {
	// 边界条件处理：空矩阵直接返回0
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	// 初始化矩阵尺寸和方向数组（上、下、左、右）
	m, n := len(matrix), len(matrix[0])
	dx := []int{-1, 1, 0, 0} // x方向移动增量
	dy := []int{0, 0, -1, 1} // y方向移动增量

	// Step 1: 构建出度矩阵（记录每个单元格有多少个更大的邻居）
	// outDegree[i][j] 表示从(i,j)出发可以到达的更大值的邻居数量
	outDegree := make([][]int, m)
	for i := range outDegree {
		outDegree[i] = make([]int, n)
	}

	// 遍历所有单元格，计算每个单元格的出度
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 检查四个方向
			for k := 0; k < 4; k++ {
				x, y := i+dx[k], j+dy[k]
				// 验证新坐标是否在矩阵范围内，且值大于当前单元格
				if x >= 0 && x < m && y >= 0 && y < n && matrix[x][y] > matrix[i][j] {
					outDegree[i][j]++ // 出度增加
				}
			}
		}
	}

	// Step 2: 初始化队列和动态规划数组
	queue := make([][]int, 0) // 用于拓扑排序的队列
	dp := make([][]int, m)    // dp[i][j] 表示以(i,j)为起点的最长路径长度
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 1 // 初始化为1，每个单元格自身构成长度为1的路径
		}
	}

	// 将出度为0的单元格（路径终点）加入队列
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if outDegree[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	maxLen := 1 // 记录全局最大路径长度

	// Step 3: 拓扑排序处理
	for len(queue) > 0 {
		cell := queue[0] // 取出队列头部单元格
		queue = queue[1:]
		i, j := cell[0], cell[1]

		// 遍历四个方向（寻找比当前单元格小的邻居）
		for k := 0; k < 4; k++ {
			x, y := i+dx[k], j+dy[k]
			// 验证新坐标是否有效，且值小于当前单元格
			if x >= 0 && x < m && y >= 0 && y < n && matrix[x][y] < matrix[i][j] {
				// 更新邻居的路径长度：当前路径长度+1 与 邻居已有路径长度 取较大值
				if dp[x][y] < dp[i][j]+1 {
					dp[x][y] = dp[i][j] + 1
					// 更新全局最大值
					if dp[x][y] > maxLen {
						maxLen = dp[x][y]
					}
				}

				// 减少邻居的出度（因为当前路径已被处理）
				outDegree[x][y]--
				// 当邻居的出度降为0时，加入队列进行处理
				if outDegree[x][y] == 0 {
					queue = append(queue, []int{x, y})
				}
			}
		}
	}

	return maxLen
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	inorder(root, &arr)
	return arr
}

func inorder(root *TreeNode, arr *[]int) {
	inorder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorder(root.Right, arr)
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println("Number of Islands:", numIslands(grid)) // 输出岛屿数量

	start := "AACCGGTT"
	end := "AACCGGTA"
	bank := []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}

	result := minMutation(start, end, bank)
	fmt.Println(result) // Output: 2

	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	result1 := longestIncreasingPath(matrix)
	fmt.Println(result1) // 输出 4
}
