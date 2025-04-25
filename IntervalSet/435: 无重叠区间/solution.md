## 435

### 题
给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。返回 需要移除区间的最小数量，使剩余区间互不重叠 。

注意 只在一点上接触的区间是 不重叠的。例如 [1, 2] 和 [2, 3] 是不重叠的。

### 题解
```go
func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {return a[1]-b[1]})
	cnt := 0  // 重叠区间计数
	preR := math.MinInt  // 前一个不重叠区间的结束位置
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
```
贪心解法，将问题改为计算最多可以选多少个互不重叠的区间，那么没选的区间就是要移除的区间。

先对区间进行排序，优先选择结束早的区间，可以留下更多空间给后续区间，从而最大化不重叠区间的数量。


