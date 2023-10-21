// Package main 2023/10/17
package main

/**
题目：寻找两个正序数组的中位数

给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

示例 1：

输入：nums1 = [1,3], nums2 = [2] 输出：2.00000 解释：合并数组 = [1,2,3] ，中位数 2
**/
func numsSort(nums1 []int, nums2 []int) []int {
	m1Len := len(nums1)
	m2Len := len(nums2)
	mid := (m1Len + m2Len) / 2
	for i := 0; i < m2Len+m1Len; i++ {

	}
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{1, 2, 3}
	numsSort(nums1, nums2)
}
