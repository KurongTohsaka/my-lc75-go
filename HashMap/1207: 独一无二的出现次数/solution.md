## 1207

### 题
给你一个整数数组 arr，如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。

### 题解
```go
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
```
用了两个哈希，第一个是对数组中所有元素进行计数，第二个是对第一个哈希元素的出现次数进行计数，目的是得到独一无二的出现次数。
