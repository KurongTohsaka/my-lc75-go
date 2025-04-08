package main

import (
	"container/heap"
	"slices"
	"sort"
)

func main() {

}

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	var ans int64

	// 边界，candidates 在后几轮不够用了
	if candidates*2+k > n {
		slices.Sort(costs)
		for _, x := range costs[:k] {
			ans += int64(x)
		}
		return ans
	}

	// 维护前后两个 candidates 的最小堆
	preHeap := hp{costs[:candidates]}
	rearHeap := hp{costs[len(costs)-candidates:]}
	heap.Init(&preHeap)
	heap.Init(&rearHeap)

	// 双指针，比较前后两个堆的最小 candidate 代价，比较 k 轮，每轮都剔除掉选择的人
	for i, j := candidates, n-1-candidates; k > 0; k-- {
		if preHeap.IntSlice[0] <= rearHeap.IntSlice[0] {
			ans += int64(preHeap.Replace(costs[i]))
			i++
		} else {
			ans += int64(rearHeap.Replace(costs[j]))
			j--
		}
	}
	return ans
}

type hp struct {
	sort.IntSlice
}

func (hp) Push(any) {}

func (hp) Pop() (_ any) { return }

func (h *hp) Replace(v int) int {
	top := h.IntSlice[0]
	h.IntSlice[0] = v
	heap.Fix(h, 0)
	return top
}
