package main

import "fmt"

func removeElements(head *ListNode, val int) *ListNode {

	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	cur := head

	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}

		cur = cur.Next
	}
	return newHead.Next

}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	list := Ints2List(a)
	fmt.Println(removeElements(list, 1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}
