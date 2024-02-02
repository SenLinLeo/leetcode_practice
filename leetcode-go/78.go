package main

import "fmt"

/**
78. 子集
https://leetcode.cn/problems/subsets/description/
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
**/
func subsets(nums []int) [][]int {
	numsLen := len(nums)
	ans := make([][]int, 0, 1<<numsLen)
	path := make([]int, 0, numsLen)

	var dfs func(int)
	dfs = func(i int) {
		if i == numsLen {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		// 不选nums[i]
		dfs(i + 1)

		// 选 nums[i]
		path = append(path, nums[i])
		dfs(i + 1)

		// 恢复现场
		path = path[:len(path)-1]
	}
	dfs(0)

	return ans
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
