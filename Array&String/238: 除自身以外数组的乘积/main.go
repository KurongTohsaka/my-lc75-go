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
