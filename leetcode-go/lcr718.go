// Package main 2024/1/21
package main

import (
	"fmt"
)

/** 718. 最长重复子数组
给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的长度最长的子数组的长度
输入：nums1=[1,2,3,2,1], nums2=[3,2,1,4,7]
输出：3
解释：长度最长的公共子数组是 [3,2,1]
**/
// dp[i][j] = dp[i - 1][j - 1] + 1
func findLength(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums1[i] == nums2[j] {
				if i-1 < 0 || j-1 < 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			}
			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
}
