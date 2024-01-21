package main

import "fmt"

// ListNode ..
type ListNode struct {
	Val  int
	Next *ListNode
}

/** 给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。
	LCR 024. 反转链表
	https://leetcode.cn/problems/UHnkqh/description/
**/
// 迭代
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// 递归
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

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

func main() {
	var head = new(ListNode)
	CreateNode(head, 10)
	PrintNode("前：", head)

	rev := reverseList(head)
	PrintNode("后：", rev)

	src := reverseList2(head)
	PrintNode("后：", src)

}
