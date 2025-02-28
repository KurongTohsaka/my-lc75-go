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
