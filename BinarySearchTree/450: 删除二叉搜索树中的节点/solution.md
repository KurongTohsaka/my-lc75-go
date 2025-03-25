## 450

### 题
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。

### 题解
```go
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root != nil {
		if root.Val < key {
			root.Right = deleteNode(root.Right, key)
		} else if root.Val > key {
			root.Left = deleteNode(root.Left, key)
		} else {
			// 这一步处理叶节点和单侧为空两种情况
			if root.Left == nil || root.Right == nil {
				if root.Left != nil {
					root = root.Left
				} else {
					root = root.Right
				}
			} else {
				// 当两侧均不为空
				// 从左子树中挑选继任者（实际左右均可）
				node := root.Left
				// 找到左子树中的最大值节点
				for node.Right != nil {
					node = node.Right
				}
				// 递归删除原左子树中的 node
				node.Left = deleteNode(root.Left, node.Val)
				node.Right = root.Right
				root = node
			}
		}
	}
	return root
}
```

删除分两种情况：

-  节点是叶子节点或只有一个子节点
   - 若左子节点存在，用左子节点替换当前节点。否则，用右子节点替换当前节点。
-  节点有两个子节点
   -  通过遍历左子树的右子节点，找到左子树中的最大值节点 node 。
   -  递归删除原左子树中的 node 。
   -  重组子树并替换当前节点。
