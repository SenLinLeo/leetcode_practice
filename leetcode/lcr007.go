// Package main 2024/1/21
package main

import (
	"fmt"
	"sort"
)

/**
LCR 007. 三数之和
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a ，b ，c 使得 a + b + c = 0 请找出所有和为 0 且 不重复 的三元组。
**/
func threeSum(nums []int) [][]int {
	numsLen := len(nums)
	ans := [][]int{}
	if numsLen < 3 {
		return ans
	}
	sort.Ints(nums)

	for i, x := range nums[:numsLen-2] {
		if i > 0 && nums[i-1] == x {
			continue
		}
		// 优化1
		if x+nums[i+1]+nums[i+2] > 0 {
			break
		}
		// 优化2
		if x+nums[numsLen-2]+nums[numsLen-1] < 0 {
			continue
		}
		left := i + 1
		right := numsLen - 1
		for left < right {
			s := x + nums[left] + nums[right]
			if s > 0 {
				right--
			} else if s < 0 {
				left++
			} else {
				ans = append(ans, []int{x, nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; left++ {
				}
				for right--; left < right && nums[right] == nums[right+1]; right-- {
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
