## 236

### 题
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

### 题解
```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 该函数的返回值代表的是 “最近公共祖先” 的候选节点，当清楚这一层含义之后，后面就好理解了。

	// 如上面所说，p、q 本身也可以作为公共祖先。递归终止条件
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 这两个的意思是，left、right 为空就代表公共祖先没在这半边子树上，就需要返回另一边不为空的作为祖先候选
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	// 当left、right 都存在，就代表候选祖先为root，或是在其两侧子树上，所以这里返回 root
	return root
}
```
注释说的很清楚了。