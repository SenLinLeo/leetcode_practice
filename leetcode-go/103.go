// Package main 2024/1/26
package main

import (
	"fmt"
)

/**
103. 二叉树的锯齿形层序遍历
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/submissions/498395821/
**/
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := [][]int{}
	cur := []*TreeNode{root}
	for flag := false; len(cur) > 0; flag = !flag {
		next := []*TreeNode{}
		vals := make([]int, len(cur))

		for i, node := range cur {
			if flag {
				vals[len(cur) - 1 - i] = node.Val
			} else {
				vals[i] = node.Val
			}
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


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	treeList := []int{1, 2, 3, 4, 5, 6, 7}

	root := createTree(treeList)
	res := zigzagLevelOrder(root)

	fmt.Println(res)
}
