## 题
给你一个整数数组 nums 和一个整数 k 。

每一步操作中，你需要从数组中选出和为 k 的两个整数，并将它们移出数组。

返回你可以对数组执行的最大操作数。

### 题解
```go
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
```
使用了哈希表来存储未匹配数的出现次数。对于每个元素找它的补数 `k-v` ，当 `hashMap[k-v] > 0` 就说明凑成了一堆数，计数器就加一。
