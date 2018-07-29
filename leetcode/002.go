package main

func main() {
	addTwoNumbers(&ListNode{4, &ListNode{3, nil}}, &ListNode{2, &ListNode{6, nil}})
	addTwoNumbers(&ListNode{5, nil}, &ListNode{5, nil})
}

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	extra := 0
	i1 := l1
	i2 := l2
	var l3, i3 *ListNode
	for ; ; {
		if i1 == nil && i2 == nil && extra == 0 {
			break
		} else if l3 == nil {
			// Initialization of result.
			l3 = &ListNode{0, nil}
			i3 = l3
		} else {
			// Initialize next node needed.
			i3.Next = &ListNode{0, nil}
			i3 = i3.Next
		}
		if i1 != nil {
			i3.Val += i1.Val
			i1 = i1.Next
		}
		if i2 != nil {
			i3.Val += i2.Val
			i2 = i2.Next
		}
		i3.Val += extra
		if i3.Val >= 10 {
			i3.Val -= 10
			extra = 1
		} else {
			extra = 0
		}
	}
	return l3
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbersV0(l1 *ListNode, l2 *ListNode) *ListNode {
	extra := 0
	i2 := l2
	i := l1
	for ; i != nil; {
		i.Val += i2.Val + extra
		if i.Val >= 10 {
			i.Val -= 10
			extra = 1
		} else {
			extra = 0
		}
		i = i.Next
		i2 = i2.Next
	}
	if extra == 1 {
		i.Next = &ListNode{1, nil}
	}
	return l1
}
