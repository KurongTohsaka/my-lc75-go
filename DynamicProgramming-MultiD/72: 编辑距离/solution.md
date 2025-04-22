## 72

### 题
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符

### 题解
```go
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	cache := make([][]int, m+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}
	for j := range n {
		cache[0][j+1] = j + 1
	}
	
	for i, c := range word1 {
		cache[i+1][0] = i + 1
		for j, c2 := range word2 {
			if c == c2 {
				cache[i+1][j+1] = cache[i][j]
			} else {
                // 分别对应删除、插入、替换操作
				cache[i+1][j+1] = min(cache[i][j+1], cache[i+1][j], cache[i][j]) + 1
			}
		} 
	}
	
	return cache[m][n]
}
```
经典难度是 Hard ，结果内卷成 Medium 了。

我们使用一个二维数组 `dp`（代码中的 `cache`），其中 `dp[i][j]` 表示将 `word1` 的前 `i` 个字符转换为 `word2` 的前 `j` 个字符所需的最少操作次数。

初始化：

- `dp[0][j]`：将空字符串转换为 `word2` 的前 `j` 个字符，需要进行 `j` 次插入操作。
- `dp[i][0]`：将 `word1` 的前 `i` 个字符转换为空字符串，需要进行 `i` 次删除操作。

对于 `dp[i][j]`，我们需要考虑 `word1[i-1]` 和 `word2[j-1]`（因为字符串索引从 0 开始）：

1. 如果 `word1[i-1] == word2[j-1]` ：
   - 当前字符相同，不需要操作，直接继承 `dp[i-1][j-1]`。
   - `dp[i][j] = dp[i-1][j-1]`
2. 如果 `word1[i-1] != word2[j-1]` ：
   - 需要进行插入、删除或替换操作：
     - 插入：在 `word1` 的第 `i` 个位置插入 `word2[j-1]`，此时 `word1` 的前 `i` 个字符与 `word2` 的前 `j-1` 个字符匹配。操作数为 `dp[i][j-1] + 1`。
     - 删除：删除 `word1` 的第 `i` 个字符，此时 `word1` 的前 `i-1` 个字符与 `word2` 的前 `j` 个字符匹配。操作数为 `dp[i-1][j] + 1`。
     - 替换：将 `word1[i-1]` 替换为 `word2[j-1]`，此时 `word1` 的前 `i-1` 个字符与 `word2` 的前 `j-1` 个字符匹配。操作数为 `dp[i-1][j-1] + 1`。
   - 取这三种操作的最小值：
     - `dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1` 
