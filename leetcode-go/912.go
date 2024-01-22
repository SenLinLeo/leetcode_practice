package main

import (
	"fmt"
	"math/rand"
)

/**
912. 排序数组
https://leetcode.cn/problems/sort-an-array/description/
**/
func sortArray(nums []int) []int {
	quickSort(nums)
	return nums
}

func quickSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// 如果你想随机选择一个索引作为基准值
	pivotIndex := rand.Intn(n)
	pivot := nums[pivotIndex]
	i, j := -1, n

	for i < j {
		for i++; nums[i] < pivot; i++ {
		}
		for j--; nums[j] > pivot; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	quickSort(nums[:j+1])
	quickSort(nums[j+1:])
}

func main() {
	list := []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54}
	listSort := sortArray(list)
	fmt.Println(listSort)
}
