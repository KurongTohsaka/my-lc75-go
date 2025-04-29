package main

func main() {

}

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
