package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	k := 5
	fmt.Println(maxOperations(nums, k))
}

func maxOperations(nums []int, k int) int {
	count := 0
	hashMap := make(map[int]int)

	for _, v := range nums {
		if hashMap[k-v] > 0 {
			hashMap[k-v]--
			count++
		} else {
			hashMap[v]++
		}
	}

	return count
}
