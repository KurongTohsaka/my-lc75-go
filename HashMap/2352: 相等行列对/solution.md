## 2352

### 题
给你一个下标从 0 开始、大小为 n x n 的整数矩阵 grid ，返回满足 Ri 行和 Cj 列相等的行列对 (Ri, Cj) 的数目。

如果行和列以相同的顺序包含相同的元素（即相等的数组），则认为二者是相等的。

### 题解
```go
func equalPairs(grid [][]int) int {
	gridHashMap := map[[200]int]int{}
	ans := 0

	n := len(grid)
	for i := 0; i < n; i++ {
		row := [200]int{}
		for j := 0; j < n; j++ {
			row[j] = grid[i][j]
		}
		gridHashMap[row]++
	}

	for j := 0; j < n; j++ {
		column := [200]int{}
		for i := 0; i < n; i++ {
			column[i] = grid[i][j]
		}
		// 这一步是寻找与这一列相同的所有行
		if gridHashMap[column] != 0 {
			ans += gridHashMap[column]
		}
	}

	return ans
}
```
先行遍历，把每一行作为键存储，然后记录相同的行出现的次数。

再列遍历，把每列作为键去查询，然后把次数加上相同的行的个数。

