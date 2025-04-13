## 17

### 题
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2021/11/09/200px-telephone-keypad2svg.png)

### 题解
```go
func letterCombinations(digits string) []string {
	res := []string{}

	if len(digits) == 0 {
		return res
	}

	// 所谓回溯，在本题实际上就是 DFS
	var backtrace func(index int, combination string)
	backtrace = func(index int, combination string) {
		// 递归跳出条件，数字已经遍历完了
		if index == len(digits) {
			res = append(res, combination)
		} else {
			letters := phoneMap[string(digits[index])]
			for _, letter := range letters {
				backtrace(index+1, combination+string(letter))
			}
		}
	}

	backtrace(0, "")

	return res
}
```
回溯解法，这里的回溯是 DFS 写法。