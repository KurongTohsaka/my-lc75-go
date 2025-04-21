## 1143


### 题
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

### 题解
```go
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	cache := make([][]int, m+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}
	// not equal: cache[i][j] = max(cache[i-1][j], cache[i][j-1])
	// equal: cache[i][j] = cache[i-1][j-1] + 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				cache[i][j] = cache[i-1][j-1] + 1
			} else {
				cache[i][j] = max(cache[i-1][j], cache[i][j-1])
			}
		}
	}

	return cache[m][n]
}
```
经典的二维DP题。

状态转移方程分为两种情况，当前字符相等时 $cache[i][j] = cache[i-1][j-1] + 1$ ，不等时 $cache[i][j] = max(cache[i-1][j], cache[i][j-1])$ 。
这两个方程还是比较好想的，相等时就是 `cache[i-1][j-1]` 再增加一个公共字符，不等时就是 `cache[i-1][j]` 和 `cache[i][j-1]` 选一个最大的。 
