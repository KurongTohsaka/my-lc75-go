## 1004
### 题
给定一个二进制数组 nums 和一个整数 k，假设最多可以翻转 k 个 0 ，则返回执行操作后 数组中连续 1 的最大个数 。

### 题解
```go
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
```
本题是变长滑动窗口，通过维护 `left`、`right` 指针来改变窗口长度，其实有些类似于快慢指针，解法与定长滑动窗口有较大不同。

窗口的长度主要由 0 的数量来控制，只要 `zeroCount <= k` 就不断移动 `right` ，直到不满足条件。此时，到移动 `left` 的时刻，
只要 `nums[left] == 0` 就减少 0 的数量且 `left` 不断向右移动，直到 0 的数量重新满足条件。最后更新最大窗口长度的值。