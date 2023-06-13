package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var (
		a ListNode
		b ListNode
	)
	a.Val = 2
	a = *a.Next
	a.Val = 4
	a = *a.Next
	a.Val = 3

	b.Val = 5
	b = *b.Next
	b.Val = 6
	b = *b.Next
	b.Val = 4

	fmt.Println(addTwoNumbers(&a, &b))
	
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		current.Val = l1.Val + l2.Val
		fmt.Println(l1.Val, "  ", l2.Val, " ", current.Val)
		l1 = l1.Next
		l2 = l2.Next
		// current=current.Next
	}

	return dummy

}
