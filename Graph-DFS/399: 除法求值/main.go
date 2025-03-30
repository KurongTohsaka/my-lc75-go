package main

func main() {

}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 构建所有变量的唯一索引
	varIndex := make(map[string]int)
	for _, eq := range equations {
		for _, v := range eq {
			if _, ok := varIndex[v]; !ok {
				varIndex[v] = len(varIndex)
			}
		}
	}

	// 初始化邻接矩阵（默认 -1）
	n := len(varIndex)
	dist := make([][]float64, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			dist[i][j] = -1.0
		}
		dist[i][i] = 1.0 // 自己除以自己是 1
	}

	// 填充已知的边
	for i, eq := range equations {
		a, b := eq[0], eq[1]
		u, v := varIndex[a], varIndex[b]
		dist[u][v] = values[i]
		dist[v][u] = 1.0 / values[i]
	}

	// Floyd-Warshall 计算所有路径
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != -1.0 && dist[k][j] != -1.0 {
					if dist[i][j] == -1.0 || dist[i][j] < dist[i][k]*dist[k][j] {
						dist[i][j] = dist[i][k] * dist[k][j]
					}
				}
			}
		}
	}

	// 处理查询
	res := make([]float64, len(queries))
	for i, q := range queries {
		a, b := q[0], q[1]
		u, hasA := varIndex[a]
		v, hasB := varIndex[b]
		if !hasA || !hasB || dist[u][v] == -1.0 {
			res[i] = -1.0
		} else {
			res[i] = dist[u][v]
		}
	}
	return res
}
