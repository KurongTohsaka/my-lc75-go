## 739

### 题
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，
下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

### 题解
```go
func dailyTemperatures(temperatures []int) []int {
	// 单调栈，栈中索引温度单调递减（从底到顶），相当于 todolist
	ms := []int{}
	ans := make([]int, len(temperatures))

	for i, t := range temperatures {
		// 循环处理单调栈内的候选项
		// 当栈内有元素，且当今天温度大于栈顶索引对应的温度时才进行处理
		for len(ms) > 0 && t > temperatures[ms[len(ms)-1]] {
			j := ms[len(ms)-1]
			// 栈顶索引出栈
			ms = ms[:len(ms)-1]
			// 更新栈顶索引那天的答案
			ans[j] = i - j
		}
		// 温度没有升高，入栈
		ms = append(ms, i)
	}

	return ans
}
```
用的单调栈，这里用的是单调递减栈。通过单调栈的特性，存储待处理的索引（温度逐渐降低）。然后碰到温度升高时，依次处理栈中元素。
