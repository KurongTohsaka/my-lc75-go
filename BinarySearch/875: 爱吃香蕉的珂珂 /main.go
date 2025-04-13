package main

import "slices"

func main() {

}

func minEatingSpeed(piles []int, h int) int {
	// 二分查找范围，注意这里的范围不是数组索引，而是香蕉数量（也是吃香蕉的速度）
	left, right := 1, slices.Max(piles)
	ans := right

	for left <= right {
		// 开始二分
		mid := left + (right-left)/2
		hours := 0
		// 统计事时间
		for _, cnt := range piles {
			// (cnt + mid - 1) / mid 等价于向上取整(cnt / mid)
			hours += (cnt + mid - 1) / mid
		}

		if hours <= h {
			ans = min(ans, mid)
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
