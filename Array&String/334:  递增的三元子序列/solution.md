## 334
### 题
给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。

如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。

```go
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
```

### 解法

本题最优解使用的是贪心算法。首先假定递增三元子序列的前两个 `first` 和 `second` ，然后通过后续的遍历不断更新这两个变量，是的后续元素更容易满足递增条件：

- 若当前元素比 `second` 大，说明已经找到，直接返回 `true` 。
- 若当前元素只比 `first` 大，则更新 `second` ，缩小后续元素比较的阈值（`second` 越小越好找第三个符合条件的值）。
- 若当前元素比 `first` 小，则更新 `first` ，尽可能让后续元素更容易找到更大的中间值。

即使 `first` 被更新为更小的值后，原先的 `second` 可能出现在 `first` 的前面（例如数组 [3,4,1,5]），但这不影响后续判断，因为只要存在一个元素比 `second` 大，即可形成三元组。