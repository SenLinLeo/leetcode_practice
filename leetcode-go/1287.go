// Package main 2024/1/22
package main

import (
	"fmt"
)

/**
1287. 有序数组中出现次数超过25%的元素
https://leetcode.cn/problems/element-appearing-more-than-25-in-sorted-array/description/
https://www.nowcoder.com/share/jump/68216111705858594671
给你一个非递减的 有序 整数数组，已知这个数组中恰好有一个整数，它的出现次数超过数组元素总数的 25%。
请你找到并返回这个整数
**/
func findSpecialInteger(arr []int) int {
	arrLen := len(arr)

	for i := 0; i < arrLen; i++ {
		j := i + 1
		for ; j < arrLen && arr[i] == arr[j]; j++ {
		}

		if j-1-i >= arrLen/4 {
			return arr[i]
		}
		i = j - 1
	}

	return 0
}

func main() {
	// 6
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
}
