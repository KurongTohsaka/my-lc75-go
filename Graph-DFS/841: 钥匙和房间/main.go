package main

func main() {

}

var (
	visited []bool
	num     int
)

func canVisitAllRooms(rooms [][]int) bool {
	num = 0
	visited = make([]bool, len(rooms))
	DFS(rooms, 0)
	return num == len(rooms)
}

func DFS(room [][]int, key int) {
	visited[key] = true
	num++

	for _, it := range room[key] {
		if !visited[it] {
			DFS(room, it)
		}
	}
}
