package main

func main() {

}

func orangesRotting(grid [][]int) int {
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(grid), len(grid[0])
	queue := [][]int{}
	fresh := 0
	time := 0

	// 初始化：收集所有腐烂橘子的位置和新鲜橘子数量
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	// 如果没有新鲜橘子，直接返回0
	if fresh == 0 {
		return 0
	}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]

			for _, d := range directions {
				x, y := cur[0]+d[0], cur[1]+d[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					fresh--
					queue = append(queue, []int{x, y})
				}
			}
		}
		if len(queue) > 0 {
			time++
		}
	}

	if fresh > 0 {
		return -1
	}

	return time
}
