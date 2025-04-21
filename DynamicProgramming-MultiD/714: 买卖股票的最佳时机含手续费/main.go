package main

func main() {

}

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	cache := make([][2]int, n)
	cache[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		// 不持有股票时，要么不买、要么买入
		cache[i][0] = max(cache[i-1][0], cache[i-1][1]+prices[i]-fee)
		// 持有股票时，要么不卖、要么卖出
		cache[i][1] = max(cache[i-1][1], cache[i-1][0]-prices[i])
	}

	return cache[n-1][0]
}
