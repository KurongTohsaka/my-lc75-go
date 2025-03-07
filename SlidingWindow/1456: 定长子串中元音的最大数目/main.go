package main

import (
	"strings"
)

func main() {

}

func maxVowels(s string, k int) int {
	maxCount := 0
	count := 0

	for i, c := range s {
		// 入
		char := string(c)
		if strings.Contains("aeiou", char) {
			count++
		}
		if i < k-1 {
			continue
		}

		// 更新
		maxCount = max(maxCount, count)

		// 出
		left := string(s[i-k+1])
		if strings.Contains("aeiou", left) {
			count--
		}
	}

	return maxCount
}
