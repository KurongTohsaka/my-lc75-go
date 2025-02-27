## 345

### 题
给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。

元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。

### 直接解法（4ms，7.7MB）
```go
package main

func main() {
	s1 := "IceCreAm"
	s2 := "leetcode"
	println(reverseVowels(s1))
	println(reverseVowels(s2))
}

func reverseVowels(s string) string {
	runes := []rune(s)
	var vowelsIdx []int

	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	for i, c := range runes {
		if vowels[c] {
			vowelsIdx = append(vowelsIdx, i)
		}
	}

	i, j := 0, len(vowelsIdx)-1
	for i < j {
		runes[vowelsIdx[i]], runes[vowelsIdx[j]] = runes[vowelsIdx[j]], runes[vowelsIdx[i]]
		i++
		j--
	}

	return string(runes)
}

```

### 双指针解法（0ms）
```go
func reverseVowels(s string) string {
	t := []byte(s)
	n := len(t)
	i, j := 0, n-1
	for i < j {
		for i < n && !strings.Contains("aeiouAEIOU", string(t[i])) {
			i++
		}
		for j > 0 && !strings.Contains("aeiouAEIOU", string(t[j])) {
			j--
		}
		if i < j {
			t[i], t[j] = t[j], t[i]
			i, j = i+1, j-1
		}
	}
	return string(t)
}
```