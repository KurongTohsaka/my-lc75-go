package main

func main() {

}

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
