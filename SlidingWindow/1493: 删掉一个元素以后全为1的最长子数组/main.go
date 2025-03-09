package main

func main() {

}

func longestSubarray(nums []int) int {
	maxLen := 0
	left := 0
	zeroCnt := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCnt++
		}

		for zeroCnt > 1 {
			if nums[left] == 0 {
				zeroCnt--
			}
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen - 1
}
