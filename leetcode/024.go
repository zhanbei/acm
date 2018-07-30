package main

// 2018-07-30T11:11:00+0800
// 2018-07-30T11:18:39+0800
// 2018-07-30T11:28:47+0800
func main() {
	swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}})
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next.Next == nil || head.Next.Next.Next == nil {
		next := head.Next
		third := next.Next
		next.Next = head
		head.Next = third
		return next
	}
	next := head.Next
	third := next.Next
	forth := third.Next
	next.Next = head
	head.Next = forth
	third.Next = swapPairs(forth.Next)
	forth.Next = third
	return next
}
