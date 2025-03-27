package main

func main() {

}

var visited []bool

func findCircleNum(isConnected [][]int) int {
	visited = make([]bool, len(isConnected))
	ans := 0

	for i, v := range visited {
		if !v {
			ans++
			DFS(isConnected, i)
		}
	}

	return ans
}

func DFS(isConnected [][]int, node int) {
	visited[node] = true

	for i, n := range isConnected[node] {
		if n == 1 && !visited[i] {
			DFS(isConnected, i)
		}
	}
}
