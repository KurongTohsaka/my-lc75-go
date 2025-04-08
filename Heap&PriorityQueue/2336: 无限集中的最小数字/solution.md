## 2336

### 题
现有一个包含所有正整数的集合 [1, 2, 3, 4, 5, ...] 。

实现 SmallestInfiniteSet 类：

SmallestInfiniteSet() 初始化 SmallestInfiniteSet 对象以包含 所有 正整数。
int popSmallest() 移除 并返回该无限集中的最小整数。
void addBack(int num) 如果正整数 num 不 存在于无限集中，则将一个 num 添加 到该无限集中。

### 题解
```go
type SmallestInfiniteSet struct {
	Elements   []int
	ElementSet map[int]bool
}

func Constructor() SmallestInfiniteSet {
	newSet := SmallestInfiniteSet{
		Elements:   make([]int, 0),
		ElementSet: make(map[int]bool),
	}
	for i := 1; i <= 1000; i++ {
		newSet.Elements = append(newSet.Elements, i)
		newSet.ElementSet[i] = true
	}
	// 构建最小堆
	for i := len(newSet.Elements)/2 - 1; i >= 0; i-- {
		newSet.MinHeapify(i)
	}
	return newSet
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	popElem := this.Elements[0]
	n := len(this.Elements)
	this.Elements[0], this.Elements[n-1] = this.Elements[n-1], this.Elements[0]
	this.Elements = this.Elements[:n-1]
	delete(this.ElementSet, popElem)
	if n-1 > 0 {
		this.MinHeapify(0)
	}
	return popElem
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if this.ElementSet[num] {
		return
	}
	this.Elements = append(this.Elements, num)
	this.ElementSet[num] = true
	// 从新添加元素的父节点开始堆化
	for i := len(this.Elements)/2 - 1; i >= 0; i-- {
		this.MinHeapify(i)
	}
}

func (this *SmallestInfiniteSet) MinHeapify(i int) {
	left := 2*i + 1
	right := 2*i + 2
	smallest := i

	if left < len(this.Elements) && this.Elements[left] < this.Elements[smallest] {
		smallest = left
	}
	if right < len(this.Elements) && this.Elements[right] < this.Elements[smallest] {
		smallest = right
	}

	if smallest != i {
		this.Elements[i], this.Elements[smallest] = this.Elements[smallest], this.Elements[i]
		this.MinHeapify(smallest)
	}
}
```
完整的构建最小堆流程，务必仔细看、背过。
