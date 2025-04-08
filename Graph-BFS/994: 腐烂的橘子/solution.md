## 994

### 题
在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：

值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。

返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。

### 题解
```go
func orangesRotting(grid [][]int) int {
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(grid), len(grid[0])
	queue := [][]int{}
	fresh := 0
	time := 0

	// 初始化：收集所有腐烂橘子的位置和新鲜橘子数量
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	// 如果没有新鲜橘子，直接返回0
	if fresh == 0 {
		return 0
	}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]

			for _, d := range directions {
				x, y := cur[0]+d[0], cur[1]+d[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					fresh--
					queue = append(queue, []int{x, y})
				}
			}
		}
		if len(queue) > 0 {
			time++
		}
	}
	
	if fresh > 0 {
		return -1
	}
	
	return time
}
```
本题是 BFS 的升级版 多源 BFS，可以看到一开始需要确定所有 BFS 开始的源点。后面的 BFS 模版略有变化，需要额外加一层遍历多源点的循环。