package main

import (
	"math"
	"slices"
)

func main() {

}

func findMinArrowShots(points [][]int) int {
	// 将气球区间按右端点升序排序，这样可以优先处理右端点较小的区间，确保箭的位置尽可能覆盖更多后续区间
	slices.SortFunc(points, func(a, b []int) int { return a[1] - b[1] })
	pre := math.MinInt
	ans := 0

	for _, point := range points {
		// 区间左侧大于箭的起始位置，不区间符合条件，需要多放一支箭
		if point[0] > pre {
			ans++
			// 把箭的位置更新为该区间的右端点（排好序了）
			pre = point[1]
		}
	}

	return ans
}
