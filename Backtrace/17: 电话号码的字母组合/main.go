package main

func main() {

}

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	res := []string{}

	if len(digits) == 0 {
		return res
	}

	// 所谓回溯，在本题实际上就是 DFS
	var backtrace func(index int, combination string)
	backtrace = func(index int, combination string) {
		// 递归跳出条件，数字已经遍历完了
		if index == len(digits) {
			res = append(res, combination)
		} else {
			letters := phoneMap[string(digits[index])]
			for _, letter := range letters {
				backtrace(index+1, combination+string(letter))
			}
		}
	}

	backtrace(0, "")

	return res
}
