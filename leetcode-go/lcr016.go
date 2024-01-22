// Package main 2023/10/21
package main

import (
	"fmt"
)

/** 给定一个字符串 s ，请你找出其中不含有重复字符的 最长连续子字符串 的长度。
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子字符串是 "abc"，所以其长度为 3。
 **/
func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	if sLen <= 1 {
		return sLen
	}

	maxLen := 1
	left := 0
	right := 0
	win := [128]int{}
	for right < sLen {
		rightChar := s[right]
		rightIndex := win[rightChar]
		if rightIndex > left {
			left = rightIndex
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}

		win[rightChar] = right + 1
		right++
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcc"))
}
