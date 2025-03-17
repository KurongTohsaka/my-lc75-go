## 206

### 题
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

### 题解
```go
func reverseList(head *ListNode) *ListNode {
	// prev 为新链表头部
	var prev, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
```
代码逻辑：
1. 初始化指针：prev 用于跟踪新链表的头节点，初始为 nil；cur 指向当前处理的节点，初始为原链表头 head。
2. 遍历链表：当 cur 非空时循环处理每个节点。
3. 反转指针：
   - 保存当前节点的下一个节点 next（避免丢失后续链表）。
   - 将当前节点的 Next 指向前驱节点 prev，实现反转。
   - prev 和 cur 同时前移：prev 移至当前节点，cur 移至之前保存的 next 节点。
4. 返回新头节点：当 cur 为空时，说明已处理完所有节点，此时 prev 即为反转后的新头节点。