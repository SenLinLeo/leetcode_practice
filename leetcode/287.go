// Package main 2023/10/21
package main

import (
	"fmt"
)

/**
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
https://leetcode.cn/problems/find-the-duplicate-number/solutions/261119/xun-zhao-zhong-fu-shu-by-leetcode-solution/?envType=study-plan-v2&envId=top-100-liked
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
	test1 := []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}
	result1 := findDuplicate(test1)
	fmt.Println(result1)
}
