## 1926

### 题
给你一个 m x n 的迷宫矩阵 maze （下标从 0 开始），矩阵中有空格子（用 '.' 表示）和墙（用 '+' 表示）。同时给你迷宫的入口 entrance ，用 entrance = [entrancerow, entrancecol] 表示你一开始所在格子的行和列。

每一步操作，你可以往 上，下，左 或者 右 移动一个格子。你不能进入墙所在的格子，你也不能离开迷宫。你的目标是找到离 entrance 最近 的出口。出口 的含义是 maze 边界 上的 空格子。entrance 格子 不算 出口。

请你返回从 entrance 到最近出口的最短路径的 步数 ，如果不存在这样的路径，请你返回 -1 。

### 题解
```go
func nearestExit(maze [][]byte, entrance []int) int {
	// 移动方向
	next := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(maze), len(maze[0])

	// 队列
	queue := [][]int{append(entrance, 0)}

	// 空地变墙，表示访问过
	maze[entrance[0]][entrance[1]] = '+'

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		curX, curY, stepCount := cur[0], cur[1], cur[2]
		for _, step := range next {
			nextX, nextY := curX+step[0], curY+step[1]
			// 下一个要移动的坐标点越界或是墙
			if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n || maze[nextX][nextY] == '+' {
				continue
			}
			// 找到出口
			if nextX == 0 || nextY == 0 || nextX == m-1 || nextY == n-1 {
				return stepCount + 1
			}
			// 变墙
			maze[nextX][nextY] = '+'
			queue = append(queue, []int{nextX, nextY, stepCount + 1})
		}
	}
	return -1
}
```
BFS 按层遍历，第一次到达出口时的步数一定是最小的。

图的 BFS 过程与之前写的树的 BFS 相似，通过队列存储访问节点，然后循环以队列长度为条件。
