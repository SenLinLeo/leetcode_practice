// Package main 2024/1/27
package main

import (
	"fmt"
)

/**
46. 全排列
https://leetcode.cn/problems/permutations/
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
**/
func permute(nums []int) [][]int {
	numsLen := len(nums)
	var ans [][]int
	path := make([]int, numsLen)
	noPath := make([]bool, numsLen)
	var dfs func(int)
	dfs = func(i int) {
		if i == numsLen {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		for j, on := range noPath {
			if !on {
				path[i] = nums[j]
				noPath[j] = true
				dfs(i + 1)
				noPath[j] = false
			}
		}
	}
	dfs(0)

	return ans
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
