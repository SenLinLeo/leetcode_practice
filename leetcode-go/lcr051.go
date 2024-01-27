// Package main 2024/1/27
package main

import (
	"math"
	"fmt"
)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
LCR 051. 二叉树中的最大路径和
https://leetcode.cn/problems/jC7MId/description/
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。
该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给定一个二叉树的根节点 root ，返回其 最大路径和，即所有路径上节点值之和的最大值。
**/
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lVal := dfs(node.Left)
		rVal := dfs(node.Right)
		ans = max(ans, lVal + rVal + node.Val)
		return max(max(lVal, rVal) + node.Val, 0)
	}

	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 3,
	}

	fmt.Println(maxPathSum(root))
}
