// Package main 2024/1/23
package main

import (
	"fmt"
)

/**
LCR 089. 打家劫舍
https://leetcode.cn/problems/Gu0c2T/description/
一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
影响小偷偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
**/
func rob(nums []int) int {
	numsLen := len(nums)
	dp := make([]int, numsLen+2)
	for i, x := range nums {
		dp[i+2] = max(dp[i+1], dp[i]+x)
	}

	return dp[numsLen+1]
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
