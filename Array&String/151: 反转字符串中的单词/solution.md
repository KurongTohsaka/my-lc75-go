## 151
### 题
给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "the sky is blue"
	s2 := "  hello world  "
	s3 := "a good   example"
	fmt.Printf("%v\n", reverseWords(s1))
	fmt.Printf("%v\n", reverseWords(s2))
	fmt.Printf("%v\n", reverseWords(s3))
}

func reverseWords(s string) string {
	var res string

	splitStrings := strings.Fields(s)

	i, j := 0, len(splitStrings)-1
	for i < j {
		splitStrings[i], splitStrings[j] = splitStrings[j], splitStrings[i]
		i++
		j--
	}

	res = strings.Join(splitStrings, " ")

	return res
}
```
先通过 Fields 函数提取出单词，然后通过双指针反转单词。