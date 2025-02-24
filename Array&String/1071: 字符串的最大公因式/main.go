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
	// 边界1, 也是最重要的条件
	if str1+str2 != str2+str1 {
		return ""
	}

	g := gcd(len(str1), len(str2))
	return str1[:g]
}

// 辗转相除法
func gcd(a int, b int) int {
	// 边界2
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
