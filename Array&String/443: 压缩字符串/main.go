package main

import "strconv"

func main() {

}

func compress(chars []byte) int {
	n := len(chars)
	write, conStart := 0, 0
	for read := 0; read < n; read++ {
		if read == n-1 || chars[read] != chars[read+1] {
			chars[write] = chars[read]
			write++
			count := read - conStart + 1
			if count > 1 {
				cntStr := strconv.Itoa(count)
				for _, c := range []byte(cntStr) {
					chars[write] = c
					write++
				}
			}
			conStart = read + 1
		}
	}
	return write
}
