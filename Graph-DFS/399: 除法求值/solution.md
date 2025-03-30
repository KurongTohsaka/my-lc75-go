## 399

### 题
给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。

另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。

返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。

注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。

注意：未在等式列表中出现的变量是未定义的，因此无法确定它们的答案。

### 题解
```go
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
```
本题难度说实话有 hard 水平了，但是是 medium 难度 😓

本题先建立每个变量对应的邻接表索引，然后初始化邻接表、用计算结果构建邻接表。

然后是本题的核心点，用 Floyd 算法计算所有可能的访问路径。最后处理所有的查询。
