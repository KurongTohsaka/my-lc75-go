package main

func main() {

}

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
