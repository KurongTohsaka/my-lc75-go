package main

func main() {

}

func onesCount(x int) (ones int) {
	// x &= x - 1 ，这一步作用是去掉最低位的 1 ，直到把所有 1 都去掉
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}
