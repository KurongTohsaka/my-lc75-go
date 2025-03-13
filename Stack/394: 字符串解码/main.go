package main

import (
	"strings"
	"unicode"
)

func main() {}

func decodeString(s string) string {
	var stack []rune

	for _, ch := range s {
		if ch != ']' {
			stack = append(stack, ch)
		} else {
			subStr := []rune{}
			for stack[len(stack)-1] != '[' {
				subStr = append(subStr, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// [ 出栈
			stack = stack[:len(stack)-1]
			reverse(subStr) // 反转顺序

			k, digit := 0, 1
			for len(stack) > 0 && unicode.IsDigit(stack[len(stack)-1]) {
				k += int(stack[len(stack)-1]-'0') * digit
				digit *= 10
				stack = stack[:len(stack)-1]
			}

			// 生成重复字符串并压入栈
			repeated := strings.Repeat(string(subStr), k)
			for _, c := range repeated {
				stack = append(stack, c)
			}
		}
	}
	return string(stack)
}

func reverse(r []rune) {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}
