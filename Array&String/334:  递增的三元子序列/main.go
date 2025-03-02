package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{5, 4, 3, 2, 1}
	nums3 := []int{2, 1, 5, 0, 4, 6}
	fmt.Println(increasingTriplet(nums1))
	fmt.Println(increasingTriplet(nums2))
	fmt.Println(increasingTriplet(nums3))
}

func increasingTriplet(nums []int) bool {
	// 边界
	if len(nums) < 3 {
		return false
	}

	first := math.MaxInt32
	second := math.MaxInt32

	for _, num := range nums {
		if num > second { // 找到第三个元素
			return true
		} else if num > first { // 更新第二个元素
			second = num
		} else { // 更新第一个元素
			first = num
		}
	}
	return false
}
