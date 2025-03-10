## 1657

### 题
如果可以使用以下操作从一个字符串得到另一个字符串，则认为两个字符串 接近 ：

操作 1：交换任意两个 现有 字符。
例如，abcde -> aecdb
操作 2：将一个 现有 字符的每次出现转换为另一个 现有 字符，并对另一个字符执行相同的操作。
例如，aacabb -> bbcbaa（所有 a 转化为 b ，而所有的 b 转换为 a ）
你可以根据需要对任意一个字符串多次使用这两种操作。

给你两个字符串，word1 和 word2 。如果 word1 和 word2 接近 ，就返回 true ；否则，返回 false 。

### 题解
```go
func closeStrings(word1 string, word2 string) bool {
	n1, n2 := len(word1), len(word2)
	if n1 != n2 {
		return false
	}

	var word1Cnt, word2Cnt [26]int

	for _, c := range word1 {
		word1Cnt[c-'a']++
	}
	for _, c := range word2 {
		word2Cnt[c-'a']++
	}

	for i := range 26 {
		if (word1Cnt[i] == 0) != (word2Cnt[i] == 0) {
			return false
		}
	}

	slices.Sort(word1Cnt[:])
	slices.Sort(word2Cnt[:])

	return word1Cnt == word2Cnt
}
```
操作 1 和 2 能实现的大前提是两个字符串的长度一致，操作 1 只要两个字符串对应的字母的计数一致就成立，操作 2 则进一步要求每个“次数”一致。

操作 1 的判断好实现，主要是操作 2 的比较需要用到排序（不排序也行，再用一个哈希表 `map[int]bool` ，用空间换时间）。

