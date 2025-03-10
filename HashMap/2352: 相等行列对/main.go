package main

func main() {

}

func equalPairs(grid [][]int) int {
	gridHashMap := map[[200]int]int{}
	ans := 0

	n := len(grid)
	for i := 0; i < n; i++ {
		row := [200]int{}
		for j := 0; j < n; j++ {
			row[j] = grid[i][j]
		}
		gridHashMap[row]++
	}

	for j := 0; j < n; j++ {
		column := [200]int{}
		for i := 0; i < n; i++ {
			column[i] = grid[i][j]
		}
		// 这一步是寻找与这一列相同的所有行
		if gridHashMap[column] != 0 {
			ans += gridHashMap[column]
		}
	}

	return ans
}
