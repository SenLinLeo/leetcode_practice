// Package main 2024/1/27
package main

import (
	"math"
	"fmt"
)

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(node *TreeNode, left, right int) bool {
	if node == nil {
		return true
	}

	tmp := node.Val
	return left < tmp && tmp < right &&
		dfs(node.Left, left, tmp) &&
		dfs(node.Right, tmp, right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	head := &TreeNode{
		Val: 2,
	}
	head.Left = &TreeNode{
		Val: 1,
	}
	head.Right = &TreeNode{
		Val: 3,
	}

	fmt.Println(isValidBST(head))
}
