## 1448

### 题
给你一棵根为 root 的二叉树，请你返回二叉树中好节点的数目。

「好节点」X 定义为：从根到该节点 X 所经过的节点中，没有任何节点的值大于 X 的值。

### 题解
```go
func goodNodes(root *TreeNode) int {
	ans := findGoodNodes(root, math.MinInt)
	return ans

}

func findGoodNodes(root *TreeNode, maxVal int) int {
	if root == nil {
		return 0
	}
	
	cnt := 0
	if root.Val >= maxVal {
		cnt++
		maxVal = root.Val
	}
	cnt += findGoodNodes(root.Left, maxVal) + findGoodNodes(root.Right, maxVal)

	return cnt
}
```
深度优先遍历，自底向上。递归返回的值为计数，关键在于把左右子树的递归遍历放在何处，这里是让计数直接与左右子树的结果相加。