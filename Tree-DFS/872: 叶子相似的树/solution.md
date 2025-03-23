## 872

### 题
请考虑一棵二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个 叶值序列 。
举个例子，如上图所示，给定一棵叶值序列为 (6, 7, 4, 9, 8) 的树。

如果有两棵二叉树的叶值序列是相同，那么我们就认为它们是 叶相似 的。

如果给定的两个根结点分别为 root1 和 root2 的树是叶相似的，则返回 true；否则返回 false 。

### 题解
```go
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	root1Order := findLeafValOrder(root1, []int{})
	root2Order := findLeafValOrder(root2, []int{})

	if len(root1Order) != len(root2Order) {
		return false
	}

	for i, val := range root1Order {
		if val != root2Order[i] {
			return false
		}
	}

	return true
}

func findLeafValOrder(root *TreeNode, order []int) []int {
	// 遍历到叶节点，返回
	if root.Left == nil && root.Right == nil {
		order = append(order, root.Val)
		return order
	}

	if root.Left != nil {
		order = findLeafValOrder(root.Left, order)
	}
	if root.Right != nil {
		order = findLeafValOrder(root.Right, order)
	}
	return order
}
```
先单独找到每棵树的叶序列，再进行对比。

要找到每棵树的叶序列，用自底向上的递归比较好。到遍历到叶节点时，将值加入序列，然后返回。当遍历完所有子树后，返回叶序列。
