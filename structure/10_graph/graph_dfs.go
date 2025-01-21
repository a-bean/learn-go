package graph

/* 深度优先遍历辅助函数 */
func dfs(g *graphAdjList, visited map[Vertex]struct{}, res *[]Vertex, vet Vertex) {
	// append 操作会返回新的的引用，必须让原引用重新赋值为新slice的引用
	*res = append(*res, vet)
	visited[vet] = struct{}{}
	// 遍历该顶点的所有邻接顶点
	for _, adjVet := range g.adjList[vet] {
		_, isExist := visited[adjVet]
		// 递归访问邻接顶点
		if !isExist {
			dfs(g, visited, res, adjVet)
		}
	}
}

/* 深度优先遍历 */
// 使用邻接表来表示图，以便获取指定顶点的所有邻接顶点
func graphDFS(g *graphAdjList, startVet Vertex) []Vertex {
	// 顶点遍历序列
	res := make([]Vertex, 0)
	// 哈希表，用于记录已被访问过的顶点
	visited := make(map[Vertex]struct{})
	dfs(g, visited, &res, startVet)
	// 返回顶点遍历序列
	return res
}
