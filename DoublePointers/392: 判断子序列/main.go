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
