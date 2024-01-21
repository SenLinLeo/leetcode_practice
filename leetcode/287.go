// Package main 2023/10/17
package main

import (
	"fmt"
)

/**
题目：287. 寻找重复数
https://leetcode.cn/problems/find-the-duplicate-number/description/
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
**/
func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow, fast := nums[0], nums[nums[0]]
	for ; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}

	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	nums1 := []int{1, 3, 3, 2}
	fmt.Println(findDuplicate(nums1))
}
