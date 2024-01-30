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

解题思路:
滑动窗口
根据题目要求，我们需要在字符串 sss 寻找字符串 ppp 的变位词。因为字符串 ppp 的变位词的长度一定与字符串 ppp 的长度相同，
所以我们可以在字符串 sss 中构造一个长度为与字符串 ppp 的长度相同的滑动窗口，并在滑动中维护窗口中每种字母的数量；
当窗口中每种字母的数量与字符串 ppp 中每种字母的数量相同时，则说明当前窗口为字符串 ppp 的变位词。
**/
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
