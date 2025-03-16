package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

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
