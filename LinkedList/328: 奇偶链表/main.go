package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

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
