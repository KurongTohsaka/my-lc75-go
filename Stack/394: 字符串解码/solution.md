## 394
### 题
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

### 题解
```go

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
```
本题要点在于只要遇到右括号就开始处理栈内元素：
- 先处理直到左括号内的所有元素，拼接后再反转（顺序是反的🐎），然后别忘记让左括号出栈。
- 之后统计出现次数，此时注意数字可能为两位数。在处理完成后，就开始按照次数重复拼接子串，然后再把该子串入栈（处理嵌套字符串的方式）。