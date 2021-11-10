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

func middleNode(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	p1 := head
	p2 := head

	for p1.Next.Next != nil {
		p1 = p1.Next.Next
		p2 = p2.Next
	}

	cursor := head
	var length int

	for cursor != nil {
		cursor = cursor.Next
		length++

	}

	if length%2 == 0 {
		return p2.Next
	}

	return p2
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

	fmt.Println(middleNode(list.head))

}
