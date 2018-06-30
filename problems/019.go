package main

// 2018-06-30T21:41:01+0800
func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 2018-06-30T21:45:24+0800
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var a []*ListNode
	node := head
	for i := 0; node != nil; i++ {
		a = append(a, node)
		node = node.Next
	}
	if len(a)-n-1 < 0 {
		return head.Next
	}
	// 10, 3
	pre := a[len(a)-n-1]
	pre.Next = pre.Next.Next
	return head
}
