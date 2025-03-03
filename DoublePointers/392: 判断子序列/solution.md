## 392
### 题
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

### 题解
```go
package main

import "fmt"

func main() {
	s1, t1 := "abc", "ahbgcd"
	s2, t2 := "axc", "ahbgdc"
	fmt.Println(isSubsequence(s1, t1))
	fmt.Println(isSubsequence(s2, t2))
}

func isSubsequence(s string, t string) bool {
	slow, fast := 0, 0
	for fast < len(t) && slow < len(s) {
		if s[slow] == t[fast] {
			slow++
		}
		fast++
	}
	return slow >= len(s)
}
```
用的双指针中的快慢指针。 因为子序列的相对顺序没有改变，使得快指针只用遍历一遍就能得到结果。

慢指针指向 s 的匹配位置，快指针指向 t 的匹配位置。 
循环的跳出条件是某一个指针超出其所指向序列的长度，最后看有没有匹配成功只需要看  slow >= len(s) 。
