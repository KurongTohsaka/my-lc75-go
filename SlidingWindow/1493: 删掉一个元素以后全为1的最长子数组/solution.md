## 1493
### 题
给你一个二进制数组 nums ，你需要从中删掉一个元素。

请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。

如果不存在这样的子数组，请返回 0 。

### 题解
```go
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
```
是变长滑动窗口，解法与 1004 高度相似。此题在于控制 0 的数量，最多为 1 。

最优想法是删掉的元素是 0 ，如若不行再删 1 ，删除元素这一步在 `maxLen - 1` 。
