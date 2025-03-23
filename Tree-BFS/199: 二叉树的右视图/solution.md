## 199

### 题
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

### 题解
```go
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
```
以下是 BFS 的基本流程：
- 声明一个空队列，用于存储待遍历的节点。
- 不管是先序遍历还是什么，大致做法都是：
  - 外层循环为 `len(queue) > 0` ，内层循环就是 `i := len(queue); i > 0; i--`，内层循环的含义是遍历队列中的所有节点（遍历完一层）。

处理当前层节点时，先右子节点后左子节点入队，保证下层队列按右到左排列，使得下次循环队首即最右节点。此时，将当前层第一个节点的值加入结果（即最右节点）。
