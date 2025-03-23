package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		queue = []*TreeNode{root}
		ans   []int
	)

	for len(queue) > 0 {
		ans = append(ans, queue[0].Val)
		for i := len(queue); i > 0; i-- {
			cur := queue[0]
			queue = queue[1:]

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
		}
	}

	return ans
}
