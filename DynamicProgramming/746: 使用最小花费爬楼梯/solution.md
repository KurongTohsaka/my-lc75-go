## 746

### 题
给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。

你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

请你计算并返回达到楼梯顶部的最低花费。

### 题解
```go
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	cache := make(map[int]int, n+1)
	cache[0] = 0
	cache[1] = 0

	for i := 2; i <= n; i++ {
		cache[i] = min(cache[i-1]+cost[i-1], cache[i-2]+cost[i-2])
	}

	return cache[n]
}
```
本题是 0-1 背包的变形，0-1 对应的两个选择是登一级或登两级阶梯。

楼梯为 `0~n-1` 级，楼梯顶对应 `n` 级。设 `cache[n]` 表示直到第 n 级阶梯的最小花费，`costs[n]` 表示从第 n 级阶梯向上爬花费。则 `cache[0], cache[1]` 用于初始化，从 `cache[2]` 开始计算：
$$
    cache[n]=min\{cache[n-1]+costs[n-1], cache[n-2]+costs[n-2]\}
$$

