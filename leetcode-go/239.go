// Package main 2023/11/5
package main

import (
	"fmt"
)

/** 239. 滑动窗口最大值
https://leetcode.cn/problems/sliding-window-maximum/description/
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
**/
func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0, len(nums)-k+1)
	var q []int

	for i, x := range nums {
		// 1. insert queue
		if len(q) > 0 && nums[q[len(q)-1]] <= x { // 当前数字 >= 队尾，及时弹出队尾（和单调栈一样）
			q = q[:len(q)-1]
		}
		q = append(q, i)

		// 2. pop queue 弹出队首不在窗口内的元素
		if i-q[0] >= k {
			q = q[1:]
		}

		// 3. record ans 记录/维护答案（根据队首）
		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}

	}

	return ans
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
