## 1456

### 题
给你字符串 s 和整数 k 。

请返回字符串 s 中长度为 k 的单个子字符串中可能包含的最大元音字母数。

英文中的 元音字母 为（a, e, i, o, u）。



### 题解
```go
func maxVowels(s string, k int) int {
	maxCount := 0  // 全局统计量
	count := 0  // 窗口统计量

	for i, c := range s {
		// 新进入
		char := string(c)
		if strings.Contains("aeiou", char) {
			count++
		}
		if i < k-1 {
			continue
		}

		// 更新统计量
		maxCount = max(maxCount, count)

		// 边界出
		left := string(s[i-k+1])
		if strings.Contains("aeiou", left) {
			count--
		}
	}

	return maxCount
}
```
本题解法是定长滑动窗口，使用通用解法：
1. 新进入：新进入的元素进行判断是否是元音字母，是的话 `count++` ，同时判断当前索引是否小于窗口。
2. 更新统计量：更新 `maxCount` 。
3. 边界出：窗口左侧边界元素 `s[i-k+1]` 是否是元音字母，是的话 `count--` 。
