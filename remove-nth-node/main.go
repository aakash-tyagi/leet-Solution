package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type linkedList struct {
	head   *ListNode
	length int
}

func (l *linkedList) prePend(n *ListNode) {
	second := l.head
	l.head = n
	l.head.Next = second
	l.length++
}

func (l linkedList) printList() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d \n", toPrint.Val)
		toPrint = toPrint.Next
		l.length--
	}

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var fast, slow *ListNode
	fast = head
	slow = head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		head = head.Next
		return head
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func main() {
	list := linkedList{}
	node1 := &ListNode{Val: 48}
	node2 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 8}

	list.prePend(node1)
	list.prePend(node2)
	list.prePend(node3)
	list.prePend(node4)

	list.printList()
	removeNthFromEnd(list.head, 1)
	list.printList()
}
