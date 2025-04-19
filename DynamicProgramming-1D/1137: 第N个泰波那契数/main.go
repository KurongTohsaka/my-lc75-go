package main

func main() {

}

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
