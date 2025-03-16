## 328

### 题
给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。

第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。

请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。

你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。

### 题解
```go
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	oddHead, evenHead := &ListNode{}, &ListNode{}
	odd, even := oddHead, evenHead

	num := 1
	for head != nil {
		if num%2 == 1 {
			odd.Next = head
			odd = odd.Next
		} else {
			even.Next = head
			even = even.Next
		}
		num++
		head = head.Next
	}

	odd.Next = evenHead.Next
	even.Next = nil
	return oddHead.Next
}
```
用两个变量记录奇偶两个链表的头节点，再用两个变量作为构建奇偶链表的指针。

然后遍历整个链表去构建奇偶链表，最后把奇链表的指针指向偶链表的头部。