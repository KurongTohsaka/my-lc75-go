## 2215

### 题
给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，请你返回一个长度为 2 的列表 answer ，其中：

answer[0] 是 nums1 中所有 不 存在于 nums2 中的 不同 整数组成的列表。
answer[1] 是 nums2 中所有 不 存在于 nums1 中的 不同 整数组成的列表。
注意：列表中的整数可以按 任意 顺序返回。



### 题解
```go
func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1Map := map[int]int{}
	nums2Map := map[int]int{}

	for _, num := range nums1 {
		nums1Map[num]++
	}
	for _, num := range nums2 {
		nums2Map[num]++
	}

	var ans1, ans2 []int

	for key, _ := range nums1Map {
		if nums2Map[key] == 0 && nums1Map[key] > 0 {
			ans1 = append(ans1, key)
		}
	}

	for key, _ := range nums2Map {
		if nums1Map[key] == 0 && nums2Map[key] > 0 {
			ans2 = append(ans2, key)
		}
	}

	return [][]int{ans1, ans2}
}
```
双哈希表记录每个元素出现次数，然后分别进行比较。
