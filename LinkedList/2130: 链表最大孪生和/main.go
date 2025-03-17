package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

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
