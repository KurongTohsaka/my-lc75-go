## 547

### 题
有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。

省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。

返回矩阵中 省份 的数量。

### 题解
```go
var visited []bool

func findCircleNum(isConnected [][]int) int {
	visited = make([]bool, len(isConnected))
	ans := 0

	for i, v := range visited {
		if !v {
			ans++
			DFS(isConnected, i)
		}
	}

	return ans
}

func DFS(isConnected [][]int, node int) {
	visited[node] = true

	for i, n := range isConnected[node] {
		if n == 1 && !visited[i] {
			DFS(isConnected, i)
		}
	}
}
```
图的深度优先遍历，以访问数组为入口进行遍历。进入遍历函数后，根据每个节点的数组继续递归。
