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
	slow := 0 // 慢指针指向待填充的位置

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
