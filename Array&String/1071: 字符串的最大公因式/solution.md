## 1071

### 题
对于字符串 `s` 和 `t`，只有在 `s = t + t + t + ... + t + t`（`t` 自身连接 1 次或多次）时，我们才认定 “`t` 能除尽 `s`”。

给定两个字符串 `str1` 和 `str2` 。返回 最长字符串 `x`，要求满足 `x` 能除尽 `str1` 且 `x` 能除尽 `str2` 。

```go
package main

func main() {
	str1 := "ABCABC"
	str2 := "ABC"
	str3 := "ABABAB"
	str4 := "ABAB"
	str5 := "LEET"
	str6 := "CODE"
	println(gcdOfStrings(str1, str2))
	println(gcdOfStrings(str3, str4))
	println(gcdOfStrings(str5, str6))
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	g := gcd(len(str1), len(str2))
	return str1[:g]
}

// 辗转相除法
func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

```

先判断两个字符串能否通过重复某个字符串构成，这个充要条件是：
- str1 + str2 == str2 + str1

如果不满足，则说明没有公共的“最大公因子”字符串。

如果满足，则令答案为 `str1` 的前 `gcd(len(str1), len(str2))` 个字符，其中 `gcd` 表示整数的最大公约数。