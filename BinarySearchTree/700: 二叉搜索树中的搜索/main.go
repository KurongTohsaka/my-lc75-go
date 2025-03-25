package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}

		if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return nil
}
