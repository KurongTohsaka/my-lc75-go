## 2542

### 题
给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，两者长度都是 n ，再给你一个正整数 k 。你必须从 nums1 中选一个长度为 k 的 子序列 对应的下标。

对于选择的下标 i0 ，i1 ，...， ik - 1 ，你的 分数 定义如下：

nums1 中下标对应元素求和，乘以 nums2 中下标对应元素的 最小值 。
用公式表示： (nums1[i0] + nums1[i1] +...+ nums1[ik - 1]) * min(nums2[i0] , nums2[i1], ... ,nums2[ik - 1]) 。
请你返回 最大 可能的分数。

一个数组的 子序列 下标是集合 {0, 1, ..., n-1} 中删除若干元素得到的剩余集合，也可以不删除任何元素。

### 题解
```go
func maxScore(nums1 []int, nums2 []int, k int) int64 {
	// 存储原始索引
	ids := make([]int, len(nums1))
	for i := range ids {
		ids[i] = i
	}

	// 根据nums2的值对索引进行降序排序
	slices.SortFunc(ids, func(i, j int) int { return nums2[j] - nums2[i] })

	// 创建一个大小为k的最小堆，并将前k个元素(nums1中对应nums2最大的k个元素)放入堆中，同时计算它们的和
	h := hp{make([]int, k)}
	sum := 0
	for i, idx := range ids[:k] {
		sum += nums1[idx]
		h.IntSlice[i] = nums1[idx]
	}
	heap.Init(&h)

	// 初始答案为前k个元素的和乘以第k个元素的nums2值(因为nums2是降序排列，所以这是前k个中最小的nums2值)
	ans := sum * nums2[ids[k-1]]
	// 遍历剩下的元素，如果当前nums1的值大于堆顶(当前堆中最小的元素)，则替换堆顶元素，更新sum，并计算新的分数
	for _, i := range ids[k:] {
		x := nums1[i]
		if x > h.IntSlice[0] {
			sum += x - h.replace(x)
			ans = max(ans, sum*nums2[i])
		}
	}
	return int64(ans)
}

// 最小堆
type hp struct{ sort.IntSlice }

func (hp) Push(any)     {}
func (hp) Pop() (_ any) { return }

// 替换堆顶元素并保持堆性质
func (h hp) replace(v int) int { top := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(&h, 0); return top }
```
本题过于复杂了，放弃了。