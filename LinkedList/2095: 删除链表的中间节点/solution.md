## 2095

### 题
给你一个链表的头节点 head 。删除 链表的 中间节点 ，并返回修改后的链表的头节点 head 。

长度为 n 链表的中间节点是从头数起第 ⌊n / 2⌋ 个节点（下标从 0 开始），其中 ⌊x⌋ 表示小于或等于 x 的最大整数。

对于 n = 1、2、3、4 和 5 的情况，中间节点的下标分别是 0、1、1、2 和 2 。

### 题解
```go
func deleteMiddle(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	slowPrev, slow, fast := dummyHead, head, head

	for fast != nil && fast.Next != nil {
		slowPrev = slow
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 删除中间节点
	slowPrev.Next = slow.Next

	return dummyHead.Next
}
```
通过快慢指针找到中间节点，同时记录慢指针的前一个节点方便后续删除操作。

本题用到了伪头节点，方便操作。
