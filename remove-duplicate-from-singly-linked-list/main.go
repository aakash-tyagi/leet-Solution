package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln ListNode) print() {
	curHead := &ln
	for curHead.Next != nil {
		fmt.Print(curHead.Val)
		curHead = curHead.Next
	}
	fmt.Println(curHead.Val)

}

func (ln *ListNode) append(val int) {
	// node last node
	newNode := &ListNode{
		Val:  val,
		Next: nil,
	}
	curHead := ln
	for curHead.Next != nil {
		curHead = curHead.Next
	}
	curHead.Next = newNode

}

func deleteDuplicates(head *ListNode) *ListNode {
	c := head
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	for c.Next != nil {
		if c.Next.Val == c.Val {
			c.Next = c.Next.Next
		} else {
			c = c.Next
		}
	}

	return head

}

func main() {
	head1 := &ListNode{}
	head1.Val = 1
	head1.append(2)
	head1.append(2)
	head1.append(4)
	head1.append(5)

	head1.print()

	deleteDuplicates(head1)

	head1.print()
}
