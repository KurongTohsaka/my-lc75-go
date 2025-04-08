package main

func main() {

}

func findKthLargest(nums []int, k int) int {
	buildMaxHeap(nums)
	heapSize := len(nums)

	// 执行k-1次弹出操作
	for i := 0; i < k-1; i++ {
		// 将堆顶元素(最大值)与最后一个元素交换
		nums[0], nums[heapSize-1] = nums[heapSize-1], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}

	// 第k次堆顶元素就是第k大的元素
	return nums[0]
}

// 构建最大堆
func buildMaxHeap(nums []int) {
	heapSize := len(nums)
	// 从最后一个非叶子节点开始向前调整
	for i := heapSize/2 - 1; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
}

// 调整堆，确保以i为根的子树满足最大堆性质
func maxHeapify(nums []int, i, heapSize int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	// 找出当前节点、左子节点、右子节点中的最大值
	if left < heapSize && nums[left] > nums[largest] {
		largest = left
	}
	if right < heapSize && nums[right] > nums[largest] {
		largest = right
	}

	// 如果最大值不是当前节点，交换并继续调整
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapify(nums, largest, heapSize)
	}
}
