## 283

### 题
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。

### 题解
```go
package main

import "fmt"

func main() {
	nums1 := []int{0, 1, 0, 3, 12}
	nums2 := []int{0}
	moveZeroes(nums1)
	moveZeroes(nums2)
	fmt.Println(nums1)
	fmt.Println(nums2)
}

func moveZeroes(nums []int) {
	slow := 0  // 慢指针指向待填充的位置

	// 快指针遍历所有非零元素
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}

	// 剩余位置补零
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}
```
该题用的双指针中的快慢指针（另一个是头尾指针）。
分为两个循环：
1. fast 先进行遍历找到非零元素，然后把非零元素复制到 slow 处。
2. 最后在 slow 处往后都是零，补零。
