// Package main 2024/1/27
package main

import (
	"fmt"
	"sort"
)

/**
217. 存在重复元素
https://leetcode.cn/problems/contains-duplicate/description/

**/
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i -1] {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1,1,1,3,3,4,3,2,4,2}) )
}

