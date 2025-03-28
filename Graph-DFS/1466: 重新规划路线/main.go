package main

func main() {

}

var (
	visited []bool
)

func minReorder(n int, connections [][]int) int {
	// 邻接表，e[i] 表示与城市 i 相连的所有边，e[i][j] 表示第 j 条边
	e := make([][][]int, n)

	/*
		对于每条原始边 [a, b] 做以下处理：
		1. 在 e[a] 中添加 [b, 1]，表示从 a 到 b 有一条原始方向的道路（1 表示需要反转）。
		2. 在 e[b] 中添加 [a, 0]，表示从 b 到 a 有一条反向的道路（0 表示不需要反转）。
	*/
	for _, edge := range connections {
		e[edge[0]] = append(e[edge[0]], []int{edge[1], 1})
		e[edge[1]] = append(e[edge[1]], []int{edge[0], 0})
	}

	return DFS(0, -1, e)
}

func DFS(curCity, visitedCity int, e [][][]int) int {
	res := 0
	for _, edge := range e[curCity] {
		// 访问过的城市的边，跳过
		if edge[0] == visitedCity {
			continue
		}

		// 如果没访问过，且路径需要反转，加入到结果中。
		res += edge[1] + DFS(edge[0], curCity, e)
	}
	return res
}
