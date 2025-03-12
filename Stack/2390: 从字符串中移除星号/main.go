package main

func main() {

}

func removeStars(s string) string {
	var stack []rune

	for _, ch := range s {
		if ch != '*' || stack[len(stack)-1] == '*' {
			stack = append(stack, ch)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}
