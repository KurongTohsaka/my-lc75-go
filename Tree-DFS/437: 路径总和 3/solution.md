## 437

### 题
给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

### 题解
```go
func pathSum(root *TreeNode, targetSum int) int {
	cnt := map[int]int{0: 1}  // 前缀和哈希表，初始0出现1次（处理根节点路径）
	var ans int
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum += root.Val  // 更新当前路径和

		// 当前前缀和与目标值的差，查看哈希表中是否存在该差值的前缀和
		ans += cnt[sum-targetSum]
		cnt[sum]++  // 将当前前缀和加入哈希表

		dfs(root.Left, sum)
		dfs(root.Right, sum)
		cnt[sum]--  // 回溯，恢复哈希表状态
	}

	dfs(root, 0)
	return ans
}
```
这道题是前缀和+回溯的形式。
在二叉树上，前缀和相当于从根节点开始的路径元素和。用哈希表 `cnt` 统计前缀和的出现次数，当我们递归到节点 `root` 时，设从根到 `root` 的路径元素和为 `sum`，
那么就找到了 `cnt[sum−targetSum]` 个符合要求的路径，加入答案。

假设从根到节点 A 的前缀和为 sumA，从根到节点 B 的前缀和为 sumB。如果 sumB - sumA = targetSum，则从节点 A 的下一个节点到节点 B 的路径和 正好是 targetSum。

```
root → ... → A → ... → B
|-----sumA-----|
|----------sumB--------|
	       |-targetSum-|  ← 这段路径的和是 sumB - sumA = targetSum
```

遍历每个节点，更新当前前缀和，并在哈希表中查找满足 `sum - targetSum` 的前缀和次数，累加到结果中。
递归返回时恢复哈希表的状态，避免不同分支的前缀和相互干扰。


