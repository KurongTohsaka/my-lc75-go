## 2130

### 题
在一个大小为 n 且 n 为 偶数 的链表中，对于 0 <= i <= (n / 2) - 1 的 i ，第 i 个节点（下标从 0 开始）的孪生节点为第 (n-1-i) 个节点 。

比方说，n = 4 那么节点 0 是节点 3 的孪生节点，节点 1 是节点 2 的孪生节点。这是长度为 n = 4 的链表中所有的孪生节点。
孪生和 定义为一个节点和它孪生节点两者值之和。

给你一个长度为偶数的链表的头节点 head ，请你返回链表的 最大孪生和 。

### 题解
```go
func pairSum(head *ListNode) int {
	// 寻找中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 反转后半部分链表
	var prev, cur *ListNode = nil, slow
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	ans := 0
	left, right := head, prev
	for right != nil {
		sum := left.Val + right.Val
		ans = max(ans, sum)
		left, right = left.Next, right.Next
	}

	return ans
}
```
本题属于缝合题，先是快慢指针找中间节点，再是反转后半部分链表，最后又是一个双指针遍历。