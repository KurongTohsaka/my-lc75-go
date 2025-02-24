## 1768

### 题
给你两个字符串 `word1` 和 `word2` 。请你从 `word1` 开始，通过交替添加字母来合并字符串。如果一个字符串比另一个字符串长，就将多出来的字母追加到合并后字符串的末尾。

返回**合并后的字符串** 。

```go
package main

func main() {
	word1 := "abc"
	word2 := "pdr"
	word3 := "ab"
	word4 := "pdrs"
	word5 := "abcd"
	word6 := "pq"
	print(mergeAlternately(word1, word2))
	print(mergeAlternately(word3, word4))
	print(mergeAlternately(word5, word6))
}

func mergeAlternately(word1 string, word2 string) string {
	var res string

	splitLen := min(len(word1), len(word2))
	for i := 0; i < splitLen; i++ {
		res += string(word1[i]) + string(word2[i])
	}

	// 边界
	if len(word1) > len(word2) {
		res += word1[splitLen:]
	} else {
		res += word2[splitLen:]
	}

	return res
}
```

这个算法采用了**交替合并**的思路，主要步骤如下：

1. **计算较短字符串的长度**：通过取两个字符串长度的较小值，确定交替合并的有效区间。
2. **遍历并交替合并字符**：用一个循环，从索引0开始同时访问两个字符串，将对应位置的字符依次添加到结果中。
3. **处理剩余部分**：当一个字符串被遍历完后，将另一个字符串剩余的部分直接追加到结果后面。

这种算法属于**线性遍历**，其时间复杂度为 O(n)，其中 n 为两个字符串长度之和，空间复杂度也是 O(n)（构造结果字符串的空间开销）。虽然有时也会将这种做法称作“双指针”方法，但实际上这里只用了一个索引变量，同时访问两个字符串。总体来说，这是一个简单直接且高效的交替合并字符串的解法。