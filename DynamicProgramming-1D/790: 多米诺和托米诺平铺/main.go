package main

func main() {

}

func numTilings(n int) int {
	if n == 1 {
		return 1
	}
	cache := make([]int, n+1)
	cache[0], cache[1], cache[2] = 1, 1, 2

	for i := 3; i <= n; i++ {
		cache[i] = (cache[i-1]*2 + cache[i-3]) % (1e9 + 7)
	}

	return cache[n]
}
