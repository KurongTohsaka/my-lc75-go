package main

func main() {

}

func nearestExit(maze [][]byte, entrance []int) int {
	// 移动方向
	next := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(maze), len(maze[0])

	// 队列
	queue := [][]int{append(entrance, 0)}

	// 空地变墙，表示访问过
	maze[entrance[0]][entrance[1]] = '+'

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		curX, curY, stepCount := cur[0], cur[1], cur[2]
		for _, step := range next {
			nextX, nextY := curX+step[0], curY+step[1]
			// 下一个要移动的坐标点越界或是墙
			if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n || maze[nextX][nextY] == '+' {
				continue
			}
			// 找到出口
			if nextX == 0 || nextY == 0 || nextX == m-1 || nextY == n-1 {
				return stepCount + 1
			}
			// 变墙
			maze[nextX][nextY] = '+'
			queue = append(queue, []int{nextX, nextY, stepCount + 1})
		}
	}
	return -1
}
