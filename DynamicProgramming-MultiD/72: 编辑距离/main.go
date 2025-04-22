package main

func main() {

}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	cache := make([][]int, m+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}
	for j := range n {
		cache[0][j+1] = j + 1
	}

	for i, c := range word1 {
		cache[i+1][0] = i + 1
		for j, c2 := range word2 {
			if c == c2 {
				cache[i+1][j+1] = cache[i][j]
			} else {
				// 分别对应删除、插入、替换操作
				cache[i+1][j+1] = min(cache[i][j+1], cache[i+1][j], cache[i][j]) + 1
			}
		}
	}

	return cache[m][n]
}
