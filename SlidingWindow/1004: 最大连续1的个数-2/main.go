package main

func main() {

}

func longestOnes(nums []int, k int) int {
	maxWinLen := 0
	zeroCount := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		// 当前窗口长度与最大长度进行对比
		maxWinLen = max(maxWinLen, right-left+1)
	}

	return maxWinLen
}
