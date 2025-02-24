package main

func main() {
	word1 := "abc"
	word2 := "pdr"
	word3 := "ab"
	word4 := "pdrs"
	word5 := "abcd"
	word6 := "pq"
	print(mergeAlternately(word1, word2))
	print(mergeAlternately(word3, word4))
	print(mergeAlternately(word5, word6))
}

func mergeAlternately(word1 string, word2 string) string {
	var res string

	splitLen := min(len(word1), len(word2))
	for i := 0; i < splitLen; i++ {
		res += string(word1[i]) + string(word2[i])
	}

	// 边界
	if len(word1) > len(word2) {
		res += word1[splitLen:]
	} else {
		res += word2[splitLen:]
	}

	return res
}
