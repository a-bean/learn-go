package main

import (
	"fmt"
)

// numIslands 计算岛屿数量
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

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println("Number of Islands:", numIslands(grid)) // 输出岛屿数量
}
