## 643

### 题
给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。

请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。

任何误差小于 10-5 的答案都将被视为正确答案。


### 题解
```go
func findMaxAverage(nums []int, k int) float64 {
	maxSum := math.MinInt
	winSum := 0

	for i, v := range nums {
		// 进入窗口
		winSum += v
		if i < k-1 {
			continue
		}

		// 更新答案
		maxSum = max(maxSum, winSum)
		// 离开窗口
		left := i - k + 1
		winSum -= nums[left]
	}
	return float64(maxSum) / float64(k)
}
```
本题属于滑动窗口中的定长滑动窗口。套路如下：

1. 入：下标为 i 的元素进入窗口，窗口元素和 s 增加 nums[i]。如果 i<k−1 则重复第一步。
2. 更新：更新答案。本题由于窗口长度固定为 k，可以统计窗口元素和的最大值 maxS，最后返回的时候再除以 k。 
3. 出：下标为 i−k+1 的窗口左边界元素要离开窗口，窗口元素和 s 减少 nums[i−k+1]。
