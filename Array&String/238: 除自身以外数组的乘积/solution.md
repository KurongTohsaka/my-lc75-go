## 238
### 题
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。

```go
package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums1))
	fmt.Println(productExceptSelf(nums2))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, len(nums))

	// 计算前缀和
	pre := make([]int, n)
	pre[0] = 1
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}

	// 计算后缀和
	suf := make([]int, n)
	suf[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i+1]
	}

	for i, p := range pre {
		res[i] = p * suf[i]
	}

	return res
}

```
### 解法

该题是一个线性动态规划问题，通过将乘积分成前缀积数组和后缀积数组，来构造状态转移方程。

状态转移方程的简单推导过程：

1. 前缀积数组（从前往后算）：
   - `pre[i]` 表示 `i` 左侧所有元素的乘积。
   - 状态转移方程：`pre[i] = nums[0] * ... * nums[i-1] = pre[i-1] * nums[i-1]` 。
   - 初始化：`pre[0] =。1` 。
2. 后缀积数组（从后往前算）：
   - `suf[i]` 表示 `i` 右侧所有元素的乘积。
   - 状态转移方程：`suf[i] = nums[i+1] * ... * nums[n-1] = suf[i+1] * nums[i+1]`。
   - 初始化：`suf[n-1] = 1`。
3. 合并结果：
   - `res[i] = pre[i] * suf[i]` 。

