## 875

### 题
珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 h 小时后回来。

珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 k 根。如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。

珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）。

### 题解
```go
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
```
本题存在某种单调性，适合用二分查找。这里的二分查找的范围注意不是索引范围，而是吃香蕉的速度，所以最小为 1 。
