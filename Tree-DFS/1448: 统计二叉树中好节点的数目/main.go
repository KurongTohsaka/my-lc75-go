package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

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
