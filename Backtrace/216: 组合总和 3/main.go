package main

func main() {

}

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	var backtrace func(int, int, []int)

	backtrace = func(start, sum int, combination []int) {
		// 递归终止
		if len(combination) == k {
			if sum == n {
				temp := make([]int, len(combination))
				copy(temp, combination)
				res = append(res, temp)
			}
			return
		}

		for i := start; i <= 9; i++ {
			if sum+i > n {
				break
			}
			backtrace(i+1, sum+i, append(combination, i))
		}
	}

	backtrace(1, 0, []int{})

	return res
}
