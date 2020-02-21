package main

import "fmt"

/**
 * Definition for singly-linked list.
*/

type ListNode struct {
      Val int
      Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	cur := head
	var pre *ListNode = nil
	var tmp *ListNode = nil

	for cur != nil {

/*	error
        pre      = cur
		cur      = cur.Next
		cur.Next = pre
*/
       // correct 1
		tmp      = cur.Next
		cur.Next = pre
		pre      = cur
		cur      = tmp

		// correct 2
		// pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要

	}
	return pre
}

func CreateNode(node *ListNode, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &ListNode{}
		cur.Next.Val = i
		cur = cur.Next
	}
}

// 打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}

func main(){
	var head = new(ListNode)
	CreateNode(head, 10)
	PrintNode("前：", head)

	yyy := reverseList(head)
	PrintNode("后：", yyy)
}