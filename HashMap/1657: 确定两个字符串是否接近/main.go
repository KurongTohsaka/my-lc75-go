package main

import "slices"

func main() {

}

func closeStrings(word1 string, word2 string) bool {
	n1, n2 := len(word1), len(word2)
	if n1 != n2 {
		return false
	}

	var word1Cnt, word2Cnt [26]int

	for _, c := range word1 {
		word1Cnt[c-'a']++
	}
	for _, c := range word2 {
		word2Cnt[c-'a']++
	}

	for i := range 26 {
		if (word1Cnt[i] == 0) != (word2Cnt[i] == 0) {
			return false
		}
	}

	slices.Sort(word1Cnt[:])
	slices.Sort(word2Cnt[:])

	return word1Cnt == word2Cnt
}
