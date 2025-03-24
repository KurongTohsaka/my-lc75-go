## 1161

### 题
给你一个二叉树的根节点 root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。

请返回层内元素之和 最大 的那几层（可能只有一层）的层号，并返回其中 最小 的那个。

### 题解
```go
func maxLevelSum(root *TreeNode) int {
	var (
		queue       []*TreeNode
		maxLayerSum = math.MinInt
		minLayer    int
		layer       int
	)
	queue = append(queue, root)

	for len(queue) > 0 {
		layer++
		tmp := 0
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			cur := queue[0]
			tmp += cur.Val
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		if tmp > maxLayerSum {
			maxLayerSum = tmp
			minLayer = layer
		}
	}
	return minLayer
}
```
标准的层序遍历。注意维护最大和、最小层两个变量即可。