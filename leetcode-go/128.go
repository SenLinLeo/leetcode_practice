// Package main 2023/10/22
package main

import (
	"fmt"
)

/**
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
时间复杂度：O(n)，其中 n 为数组的长度。具体分析已在上面正文中给出。
空间复杂度：O(n))。哈希表存储数组中所有的数需要 O(n)O(n)O(n) 的空间
**/
func longestConsecutive(nums []int) int {
	numsMap := make(map[int]bool)
	for num := range nums {
		numsMap[num] = true
	}

	longSteak := 0
	for num := range nums {
		if !numsMap[num-1] {
			currentNum := num
			currentSteak := 1
			for numsMap[currentNum+1] {
				currentNum++
				currentSteak++
			}
			if longSteak < currentSteak {
				longSteak = currentSteak
			}
		}
	}

	return longSteak
}

func main() {
	ret := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	fmt.Println(ret)
}
