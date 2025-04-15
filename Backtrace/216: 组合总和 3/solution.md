## 216

### 题
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：

只使用数字1到9
每个数字 最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。

### 题解
```go
func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	var backtrace func(int, int, []int)

	backtrace = func(start, sum int, combination []int) {
		// 递归终止
		if len(combination) == k {
			if sum == n {
				temp := make([]int, len(combination))
				copy(temp, combination)
				res = append(res, temp)
			}
			return
		}

		for i := start; i <= 9; i++ {
			if sum+i > n {
				break
			}
			backtrace(i+1, sum+i, append(combination, i))
		}
	}

	backtrace(1, 0, []int{})

	return res
}
```
回溯解法，还是 DFS 写法。
这里需要注意当递归终止条件触发时，加入的应该是 combination 的复制（因为该变量还需要变更）。
每次递归都要从 start 开始依次遍历，这里有一个剪枝条件 `sum+i > n` 。