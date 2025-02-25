package main

func main() {
	candies1 := []int{2, 3, 5, 1, 3}
	candies2 := []int{4, 2, 1, 1, 2}
	candies3 := []int{12, 1, 12}
	extraCandies1 := 3
	extraCandies2 := 1
	extraCandies3 := 10
	println(kidsWithCandies(candies1, extraCandies1))
	println(kidsWithCandies(candies2, extraCandies2))
	println(kidsWithCandies(candies3, extraCandies3))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var result []bool
	var maxVal int
	for i := 0; i < len(candies); i++ {
		maxVal = max(maxVal, candies[i])
	}

	for i, _ := range candies {
		if (candies[i] + extraCandies) >= maxVal {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}
