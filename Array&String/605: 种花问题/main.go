package main

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n1 := 1
	n2 := 2
	println(canPlaceFlowers(flowerbed, n1))
	println(canPlaceFlowers(flowerbed, n2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)

	// 边界
	if length == 1 {
		if flowerbed[0] == 0 {
			n -= 1
		}
	} else {
		for i := 0; i < length; i++ {
			if flowerbed[i] == 0 && ((i == 0 && flowerbed[i+1] == 0) || (i == length-1 && flowerbed[i-1] == 0) || (i > 0 && i < length-1 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0)) {
				flowerbed[i] = 1
				n -= 1
			}
		}
	}
	return n <= 0
}
