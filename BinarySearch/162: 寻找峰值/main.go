package main

func main() {

}

func findPeakElement(nums []int) int {
	left, right := -1, len(nums)-1 // 开区间 (-1, n-1)
	// left+1 < right 保证了区间内至少有两个元素
	// 特殊的二分查找
	for left+1 < right { // 开区间不为空
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
