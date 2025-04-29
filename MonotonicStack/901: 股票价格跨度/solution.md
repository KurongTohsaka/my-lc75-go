## 901

### 题
设计一个算法收集某些股票的每日报价，并返回该股票当日价格的 跨度 。

当日股票价格的 跨度 被定义为股票价格小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）。

例如，如果未来 7 天股票的价格是 [100,80,60,70,60,75,85]，那么股票跨度将是 [1,1,1,2,1,4,6] 。

实现 StockSpanner 类：

StockSpanner() 初始化类对象。
int next(int price) 给出今天的股价 price ，返回该股票当日价格的 跨度 。


### 题解
```go
type StockSpanner struct {
	stack [][2]int // 存储价格和对应的跨度
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	span := 1

	// 弹出栈顶所有小于等于当前价格的元素，累加它们的跨度
	for len(this.stack) > 0 && this.stack[len(this.stack)-1][0] <= price {
		span += this.stack[len(this.stack)-1][1]
		this.stack = this.stack[:len(this.stack)-1]
	}

	this.stack = append(this.stack, [2]int{price, span})
	return span
}
```
题目的意思是找上一个大于它的元素的位置，那么就必须存储连续个比它小的元素才能计算跨度。

本题也是单调栈的经典应用，本题是单调递减栈。

