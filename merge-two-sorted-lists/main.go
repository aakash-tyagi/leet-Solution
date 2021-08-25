package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	if l1 == nil && l2 == nil {
		return nil
	}

	l3 := &ListNode{}

	if l1.Val < l2.Val {
		l3.Val = l1.Val
		l1 = l1.Next
		l3.Next = mergeTwoSortedLists(l1, l2)
	} else {
		l3.Val = l2.Val
		l2 = l2.Next
		l3.Next = mergeTwoSortedLists(l1, l2)
	}

	return l3

}

func main() {

}
