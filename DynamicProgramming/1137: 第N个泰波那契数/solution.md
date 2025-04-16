## 1137

### 题
泰波那契序列 Tn 定义如下：

T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2

给你整数 n，请返回第 n 个泰波那契数 Tn 的值。
 
### 题解
```go
func tribonacci(n int) int {
	cache := map[int]int{}
	cache[0] = 0
	cache[1] = 1
	cache[2] = 1

	v, ok := cache[n]
	if ok {
		return v
	}

	// 计算
	for i := 3; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2] + cache[i-3]
	}

	return cache[n]
}
```
前缀和，动态规划的前置步骤。