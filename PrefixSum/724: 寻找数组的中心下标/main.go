package main

func main() {

}

func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	leftSum := 0
	for i, v := range nums {
		if leftSum*2+v == sum {
			return i
		}
		leftSum += v
	}
	return -1
}
