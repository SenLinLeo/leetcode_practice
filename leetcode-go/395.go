package main

import (
	"fmt"
	"strings"
)

/**
395. 至少有 K 个重复字符的最长子串
https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/description/
给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。
如果不存在这样的子字符串，则返回 0。

示例 1：
输入：s = "aaabb", k = 3
输出：3
解释：最长子串为 "aaa" ，其中 'a' 重复了 3 次。
**/
func longestSubstring(s string, k int) int {
	sLen := len(s)
	cnt := [26]int{}
	//记录每个字符出现的次数
	for i := 0; i < sLen; i++ {
		cnt[s[i]-'a']++
	}
	var split byte
	for i := 0; i < sLen; i++ {
		//如果出现次数小于k次，那么结果字符串中肯定不包括该字符，那就将s分割
		if cnt[s[i]-'a'] < k {
			split = s[i]
		}
	}
	if split == 0 {
		return sLen
	}
	newSplit := string(split)
	ans := 0
	for _, subStr := range strings.Split(s, newSplit) { //对s分割结果重复执行函数
		ans = max(ans, longestSubstring(subStr, k))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println()
}
