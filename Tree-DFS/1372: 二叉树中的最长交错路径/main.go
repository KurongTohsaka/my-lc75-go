package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

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
