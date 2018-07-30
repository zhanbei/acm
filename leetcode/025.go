package main

import "fmt"

// 2018-07-30T11:46:09+0800
// 2018-07-30T11:59:35+0800
func main() {
	res := reverseKGroup(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2)
	fmt.Println(res, res.Next, res.Next.Next, res.Next.Next.Next)
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	next := head
	for i := 0; i < k; i++ {
		if next == nil {
			return head
		}
		next = next.Next
	}
	tail := next
	next = head
	// Start reverse the k continuous nodes.
	nodes := make([]*ListNode, k)
	for i := 0; i < k; i++ {
		nodes[i] = next
		next = next.Next
	}
	for i := k - 1; i > 0; i-- {
		nodes[i].Next = nodes[i-1]
	}
	nodes[0].Next = reverseKGroup(tail, k)
	return nodes[k-1]
}
