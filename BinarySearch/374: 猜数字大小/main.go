package main

func main() {

}

func guess(num int) int

func guessNumber(n int) int {
	for low := 1; low < n; {
		mid := low + (n-low)/2
		if guess(mid) == 1 {
			low = mid + 1
		} else {
			n = mid
		}
	}
	return n
}
