## 11
###  题
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。



### 题解
```go
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxCap := 0

	for left < right {
		if height[left] < height[right] {
			capacity := height[left] * (right - left)
			if capacity > maxCap {
				maxCap = capacity
			}
			left++
		} else {
			capacity := height[right] * (right - left)
			if capacity > maxCap {
				maxCap = capacity
			}
			right--
		}
	}
	return maxCap
}
```
经典接雨水问题的简单版，用双指针中的首尾指针解决。通过首尾指针模拟两块板子，最大容量由最短板决定，所以指针的移动条件就是某个指针指向的板较小。

