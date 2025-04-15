package unionfindset

// UnionFind 并查集接口
type UnionFind interface {
	// Union 合并两个元素所在的集合
	Union(p, q int)
	// Find 查找元素所属的集合
	Find(p int) int
	// Connected 判断两个元素是否属于同一个集合
	Connected(p, q int) bool
	// Count 返回集合的数量
	Count() int
}

// unionFind 并查集实现
type unionFind struct {
	parent []int // 存储节点的父节点
	rank   []int // 基于rank的优化，记录树的高度
	count  int   // 连通分量的数量
}

// NewUnionFind 创建并查集
func NewUnionFind(n int) UnionFind {
	uf := &unionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		count:  n,
	}

	// 初始化，每个节点的父节点都是自己
	for i := 0; i < n; i++ {
		uf.parent[i] = i // 设置i(左)的父节点i
		uf.rank[i] = 1
	}

	return uf
}

// Find 查找元素所属的集合（路径压缩优化）:找到p的祖先
func (uf *unionFind) Find(p int) int {
	if p != uf.parent[p] {
		// 路径压缩：将节点直接连接到根节点
		uf.parent[p] = uf.Find(uf.parent[p])
	}
	return uf.parent[p]
}

// Union 合并两个元素所在的集合（基于rank的优化）
func (uf *unionFind) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)

	if rootP == rootQ {
		return
	}

	// 将rank小的树连接到rank大的树上
	if uf.rank[rootP] < uf.rank[rootQ] {
		uf.parent[rootP] = rootQ
	} else if uf.rank[rootP] > uf.rank[rootQ] {
		uf.parent[rootQ] = rootP
	} else {
		uf.parent[rootQ] = rootP
		uf.rank[rootP]++
	}

	uf.count--
}

// Connected 判断两个元素是否属于同一个集合
func (uf *unionFind) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Count 返回集合的数量
func (uf *unionFind) Count() int {
	return uf.count
}
