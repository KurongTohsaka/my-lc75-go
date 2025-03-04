package main

func main() {

}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxCap := 0

	for left < right {
		if height[left] < height[right] {
			capacity := height[left] * (right - left)
			if capacity > maxCap {
				maxCap = capacity
			}
			left++
		} else {
			capacity := height[right] * (right - left)
			if capacity > maxCap {
				maxCap = capacity
			}
			right--
		}
	}
	return maxCap
}
