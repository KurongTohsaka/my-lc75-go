## 1372

### 题
给你一棵以 root 为根的二叉树，二叉树中的交错路径定义如下：

选择二叉树中 任意 节点和一个方向（左或者右）。
如果前进方向为右，那么移动到当前节点的的右子节点，否则移动到它的左子节点。
改变前进方向：左变右或者右变左。
重复第二步和第三步，直到你在树中无法继续移动。
交错路径的长度定义为：访问过的节点数目 - 1（单个节点的路径长度为 0 ）。

请你返回给定树中最长 交错路径 的长度。



### 题解
```go
func longestZigZag(root *TreeNode) int {
	var ans int
	var dfs func(*TreeNode, int, string)

	dfs = func(root *TreeNode, cnt int, direction string) {
		if root == nil {
			return
		}

		ans = max(ans, cnt)
		
		if root.Left != nil {
			// 需要保持下一个递归的方向与当前不同
			// 如果相同，那就以当前节点为新起点
			if direction != "l" {
				dfs(root.Left, cnt+1, "l")
			} else {
				dfs(root.Left, 1, direction)
			}
		}

		if root.Right != nil {
			if direction != "r" {
				dfs(root.Right, cnt+1, "r")
			} else {
				dfs(root.Right, 1, direction)
			}
		}
	}

	dfs(root, 0, "")

	return ans
}
```
zigzag 锯齿状路径，必须满足每个相邻节点的“方向”不同。所以在处理递归时，需要额外维护 `direction` 前进方向。进入下一个递归时，方向必须保持与当前遍历方向不同，
如果相同，就以当前节点为新起点进行下一次递归。