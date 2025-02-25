## 1431

### 题
有 n 个有糖果的孩子。给你一个数组 candies，其中 candies[i] 代表第 i 个孩子拥有的糖果数目，和一个整数 extraCandies 表示你所有的额外糖果的数量。

返回一个长度为 n 的布尔数组 result，如果把所有的 extraCandies 给第 i 个孩子之后，他会拥有所有孩子中 最多 的糖果，那么 result[i] 为 true，否则为 false。

注意，允许有多个孩子同时拥有 最多 的糖果数目。

```go

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

```

先找最大值，不用单独再排序找，这样空间复杂度就下来了。