// Package main 2024/1/23
package main

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

