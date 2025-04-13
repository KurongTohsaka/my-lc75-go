## 162

### 题
峰值元素是指其值严格大于左右相邻值的元素。

给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

你可以假设 nums[-1] = nums[n] = -∞ 。

你必须实现时间复杂度为 O(log n) 的算法来解决此问题。

### 题解
```go
func findPeakElement(nums []int) int {
	left, right := -1, len(nums)-1 // 开区间 (-1, n-1)
	// left+1 < right 保证了区间内至少有两个元素
	// 特殊的二分查找
	for left+1 < right {           // 开区间不为空
		mid := left + (right-left)/2
		// 比较 nums[mid] 和 nums[mid+1] 就能确定峰值的可能方向
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
```
只需要比较 nums[mid] 和 nums[mid+1] 就能确定峰值的可能方向。利用二分查找的思想，通过比较局部信息来缩小搜索范围，无需有序数组。
这是二分查找的一种变体，适用于具有特定单调性或局部极值的问题。