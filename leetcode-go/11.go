// Package main 2024/1/21
package main

import (
	"fmt"
)

/**
11. 盛最多水的容器
https://leetcode.cn/problems/container-with-most-water/description/
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
**/
func maxArea(height []int) int {
	ans := 0
	left := 0
	right := len(height) - 1

	for left < right {
		area := (right - left) * min(height[right], height[left])
		if area > ans {
			ans = area
		}
		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	// 49
	fmt.Println(maxArea([]int{1, 7, 3, 2, 4, 5, 8, 2, 7}))
}
