package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

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
