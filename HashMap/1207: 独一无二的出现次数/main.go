package main

func main() {

}

func uniqueOccurrences(arr []int) bool {
	cntHash := map[int]int{}
	uniqueCntHash := map[int]int{}

	for _, v := range arr {
		cntHash[v]++
	}

	for _, v := range cntHash {
		uniqueCntHash[v]++
	}

	for _, v := range uniqueCntHash {
		if v >= 2 {
			return false
		}
	}

	return true
}
