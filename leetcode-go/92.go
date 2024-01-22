// Package main 2024/1/22
package main

// 92. 反转链表 II https://leetcode.cn/problems/reverse-linked-list-ii/description/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	p0 := dummy

	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}

	var pre, cur *ListNode = nil, p0.Next
	for i := 0; i < right-left+1; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	p0.Next.Next = cur
	p0.Next = pre
	return dummy.Next
}

/**
func CreateNode(node *ListNode, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &ListNode{}
		cur.Next.Val = i
		cur = cur.Next
	}
}

// PrintNode 打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}
**/

func main() {
	var head = new(ListNode)
	CreateNode(head, 10)
	rev := reverseBetween(head, 2, 3)

	PrintNode("=", rev)
}
