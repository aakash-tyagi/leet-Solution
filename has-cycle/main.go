package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	fast := head
	slow := head

	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}

	}

	return false

}

func main() {
	a := []int{1, 2, 3, 4}
	list1 := Ints2ListWithCycle(a, 2)
	list2 := Ints2List(a)

	fmt.Println(hasCycle(list1))
	fmt.Println(hasCycle(list2))

}

//-----------structure for cycled linked list----------------
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

func Ints2ListWithCycle(nums []int, pos int) *ListNode {
	head := Ints2List(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}
