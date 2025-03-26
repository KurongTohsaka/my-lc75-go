## 841

### 题
有 n 个房间，房间按从 0 到 n - 1 编号。最初，除 0 号房间外的其余所有房间都被锁住。你的目标是进入所有的房间。然而，你不能在没有获得钥匙的时候进入锁住的房间。

当你进入一个房间，你可能会在里面找到一套 不同的钥匙，每把钥匙上都有对应的房间号，即表示钥匙可以打开的房间。你可以拿上所有钥匙去解锁其他房间。

给你一个数组 rooms 其中 rooms[i] 是你进入 i 号房间可以获得的钥匙集合。如果能进入 所有 房间返回 true，否则返回 false。

### 题解
```go
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
```
图的深度优先遍历。必须要维护一个访问数组，记录节点访问情况。这里的深度优先遍历的参数是图本身和要访问的节点。

当一个节点记录完成，就把相应节点中存在的值依次输入到递归函数中。