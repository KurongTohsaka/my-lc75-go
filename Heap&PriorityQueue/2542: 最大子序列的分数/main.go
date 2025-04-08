package main

import (
	"container/heap"
	"slices"
	"sort"
)

func main() {

}

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
