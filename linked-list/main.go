package main

import "fmt"

type node struct {
	val  int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prePend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printList() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d \n", toPrint.val)
		toPrint = toPrint.next
		l.length--
	}

}

func main() {
	list := linkedList{}
	node1 := &node{val: 48}
	node2 := &node{val: 4}
	node3 := &node{val: 3}
	node4 := &node{val: 8}

	list.prePend(node1)
	list.prePend(node2)
	list.prePend(node3)
	list.prePend(node4)

	list.printList()
}
