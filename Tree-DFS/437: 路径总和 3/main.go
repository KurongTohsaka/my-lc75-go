package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func pathSum(root *TreeNode, targetSum int) int {
	cnt := map[int]int{0: 1}
	var ans int
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}

		sum += root.Val           // 更新前缀和
		ans += cnt[sum-targetSum] // 找前缀和为 sum-targetSum 的路径是否存在，不等于 0 就代表这一段路径存在，就说明该段路径的下一个节点开始到本节点的路径和就是 targetSum
		cnt[sum]++                // 把该前缀和对应计数增加

		dfs(root.Left, sum)
		dfs(root.Right, sum)

		cnt[sum]-- // 回溯
	}

	dfs(root, 0)
	return ans
}
