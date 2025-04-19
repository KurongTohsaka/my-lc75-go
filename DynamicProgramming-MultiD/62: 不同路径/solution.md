## 62

### 题
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

### 题解
```go
func uniquePaths(m int, n int) int {
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}
	cache[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// i = 0, 只能往右走
			if i == 0 && j > 0 {
				cache[i][j] = cache[i][j-1]
			}
			// j = 0, 只能向下走
			if j == 0 && i > 0 {
				cache[i][j] = cache[i-1][j]
			}
			if j > 0 && i > 0 {
				cache[i][j] = cache[i-1][j] + cache[i][j-1]
			}
		}
	}

	return cache[m-1][n-1]
}
```
状态转移方程分为三种，一个是在第一行只能向右走，一个是在第一列只能向下走，另一个则是向右向下都可以。

分别对应的状态转移方程如下：$cache[i][j] = cache[i][j-1]$ 、 $cache[i][j] = cache[i-1][j]$ 、 $cache[i][j] = cache[i-1][j] + cache[i][j-1]$ .

