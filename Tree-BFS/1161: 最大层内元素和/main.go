package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

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
