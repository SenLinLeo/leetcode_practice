// Package main 2023/10/8
package main

import (
	"fmt"
)

/**
LCR 015. 找到字符串中所有字母异位词
https://leetcode.cn/problems/VabMRr/description/
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引

示例 1：
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的变位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的变位词。
**/
// 滑动窗口
func findAnagrams(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}
	const CharNum = 26
	var sCount, pCount [CharNum]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
