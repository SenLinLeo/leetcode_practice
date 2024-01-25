// Package main 2024/1/23
package main

import (
	"fmt"
)

/** 102. 二叉树的层序遍历
https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
**/
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	cur := []*TreeNode{root}
	for len(cur) > 0 {
		next := []*TreeNode{}
		vals := make([]int, len(cur))

		for i, node := range cur {
			vals[i] = node.Val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		cur = next
		ans = append(ans, vals)
	}

	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTree(list []int) *TreeNode {
	if len(list) == 0 {
		return nil
	}

	root := &TreeNode{Val: list[0]}
	queue := []*TreeNode{root}

	i := 1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if i < len(list) {
			node.Left = &TreeNode{Val: list[i]}
			queue = append(queue, node.Left)
			i++
		}

		if i < len(list) {
			node.Right = &TreeNode{Val: list[i]}
			queue = append(queue, node.Right)
			i++
		}
	}

	return root
}

func main() {
	treeList := []int{1, 2, 3, 4, 5, 6, 7}

	root := createTree(treeList)
	// root.Left = &TreeNode{Val: 2}
	// root.Right = &TreeNode{Val: 3}
	// root.Left.Left = &TreeNode{Val: 4}
	// root.Left.Right = &TreeNode{Val: 5}
	// root.Right.Left = &TreeNode{Val: 6}
	// root.Right.Right = &TreeNode{Val: 7}
	res := levelOrder(root)

	fmt.Println(res)
}
