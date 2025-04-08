## 2462

### 题
给你一个下标从 0 开始的整数数组 costs ，其中 costs[i] 是雇佣第 i 位工人的代价。

同时给你两个整数 k 和 candidates 。我们想根据以下规则恰好雇佣 k 位工人：

总共进行 k 轮雇佣，且每一轮恰好雇佣一位工人。
在每一轮雇佣中，从最前面 candidates 和最后面 candidates 人中选出代价最小的一位工人，如果有多位代价相同且最小的工人，选择下标更小的一位工人。
比方说，costs = [3,2,7,7,1,2] 且 candidates = 2 ，第一轮雇佣中，我们选择第 4 位工人，因为他的代价最小 [3,2,7,7,1,2] 。
第二轮雇佣，我们选择第 1 位工人，因为他们的代价与第 4 位工人一样都是最小代价，而且下标更小，[3,2,7,7,2] 。注意每一轮雇佣后，剩余工人的下标可能会发生变化。
如果剩余员工数目不足 candidates 人，那么下一轮雇佣他们中代价最小的一人，如果有多位代价相同且最小的工人，选择下标更小的一位工人。
一位工人只能被选择一次。
返回雇佣恰好 k 位工人的总代价。

### 题解
```go
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
```
两个最小堆+双指针，难度有的。最小堆使用标准库实现。