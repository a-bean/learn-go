package main

// 684: 冗余连接 https://leetcode.cn/problems/redundant-connection/description/
// TODO
// 207: 课程表 https://leetcode.cn/problems/course-schedule/description/
/*
AOV 网的拓扑排序
 AOV网（Activity On Vertex NetWork）用顶点表示活动，边表示活动（顶点）发生的先后关系。AOV网的边不设权值，若存在边<a,b>则表示活动a必须发生在活动b之前。
 若网中所有活动均可以排出先后顺序（任两个活动之间均确定先后顺序），则称网是拓扑有序的，这个顺序称为网上一个全序。(详情参见离散数学/图论相关内容)。
 在AOV网上建立全序的过程称为拓扑排序的过程，这个算法并不复杂：
 	1. 在网中选择一个入度为0的顶点输出
 	2. 在图中删除该顶点及所有以该顶点为尾的边
 	3. 重复上述过程，直至所有边均被输出。
 若图中无入度为0的点未输出，则图中必有环。
*/

func canFinish(n int, pre [][]int) bool {
	// in 数组记录每个课程的入度（即有多少个课程需要先修）
	in := make([]int, n)
	// frees 数组表示每个课程的后续课程（哪些课程可以在该课程之后修）
	frees := make([][]int, n)
	// next 队列，用来存储当前没有前置课程限制的课程
	next := make([]int, 0, n)

	// 遍历所有先修课程的关系
	for _, v := range pre {
		// v[0] 课程需要增加入度，v[1] 课程的后续课程列表加入 v[0]
		in[v[0]]++
		frees[v[1]] = append(frees[v[1]], v[0])
	}

	// 将入度为 0 的课程添加到 next 队列中，这些课程可以立即修
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			next = append(next, i)
		}
	}

	// 开始处理队列中的课程
	for i := 0; i != len(next); i++ {
		c := next[i]  // 取出队列中的一个课程
		v := frees[c] // 获取该课程的后续课程列表
		// 遍历后续课程
		for _, vv := range v {
			// 减少后续课程的入度
			in[vv]--
			// 如果该课程的入度为 0，说明它可以被修，可以加入 next 队列
			if in[vv] == 0 {
				next = append(next, vv)
			}
		}
	}

	return len(next) == n
}

// 210: 课程表 II https://leetcode.cn/problems/course-schedule-ii/description/
func findOrder(numCourses int, prerequisites [][]int) []int {
	in := make([]int, numCourses)
	frees := make([][]int, numCourses)
	next := make([]int, 0, numCourses)
	for _, v := range prerequisites {
		in[v[0]]++
		frees[v[1]] = append(frees[v[1]], v[0])
	}
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			next = append(next, i)
		}
	}
	for i := 0; i != len(next); i++ {
		c := next[i]
		v := frees[c]
		for _, vv := range v {
			in[vv]--
			if in[vv] == 0 {
				next = append(next, vv)
			}
		}
	}
	if len(next) == numCourses {
		return next
	}
	return []int{}
}

func main() {
	canFinish(3, [][]int{{0, 1}, {0, 2}, {1, 2}})
	findOrder(3, [][]int{{0, 1}, {0, 2}, {1, 2}})
}
