package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

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
