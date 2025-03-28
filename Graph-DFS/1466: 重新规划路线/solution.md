## 1466

### 题
n 座城市，从 0 到 n-1 编号，其间共有 n-1 条路线。因此，要想在两座不同城市之间旅行只有唯一一条路线可供选择（路线网形成一颗树）。去年，交通运输部决定重新规划路线，以改变交通拥堵的状况。

路线用 connections 表示，其中 connections[i] = [a, b] 表示从城市 a 到 b 的一条有向路线。

今年，城市 0 将会举办一场大型比赛，很多游客都想前往城市 0 。

请你帮助重新规划路线方向，使每个城市都可以访问城市 0 。返回需要变更方向的最小路线数。

题目数据 保证 每个城市在重新规划路线方向后都能到达城市 0 。

### 题解
```go
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

		// 如果没访问过，且路径需要反转（原方向），加入到结果中。
		res += edge[1] + DFS(edge[0], curCity, e)
	}
	return res
}
```
题目可以转化为“从城市 0 出发，所有道路的方向必须指向 0 或其子节点”。因此，我们需要统计所有“背离 0”的道路数量。

将所有路径保存到邻接表中，然后开始 DFS：
- 从城市 0 出发，沿着所有可能的道路遍历。
- 如果遇到原始方向的道路（edge[1] == 1），说明这条道路的方向是背离 0 的，需要反转。
- 递归累加所有需要反转的道路数量。

