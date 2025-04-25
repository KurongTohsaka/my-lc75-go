package main

import (
	"math"
	"slices"
)

func main() {

}

func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int { return a[1] - b[1] })
	cnt := 0            // 重叠区间计数
	preR := math.MinInt // 前一个不重叠区间的结束位置
	for _, p := range intervals {
		// 当前区间的开始 >= 前一个区间的结束
		if p[0] >= preR {
			// 重叠区间 + 1
			cnt++
			preR = p[1]
		}
	}

	return len(intervals) - cnt
}
