package main

import "math"

func main() {

}

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
