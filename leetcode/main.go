// Package main 2023/10/8
package main

import (
	"fmt"
	"reflect"
	"sort"
)

// 给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引
func findAnagramsN(s string, p string) []int {
	res := make([]int, 0)
	pCount := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		pCount[p[i]]++
	}
	sCount := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sCount[s[i]]++
		if i >= len(p) {
			sCount[s[i-len(p)]]--
			if sCount[s[i-len(p)]] == 0 {
				delete(sCount, s[i-len(p)])
			}
		}
		if reflect.DeepEqual(sCount, pCount) {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

// findAnagramsLogN 可以使用滑动窗口和排序来实现
func findAnagramsLogN(s string, p string) []int {
	res := make([]int, 0)
	pSort := sortString(p)
	for i := 0; i <= len(s)-len(p); i++ {
		sSort := sortString(s[i : i+len(p)])
		if sSort == pSort {
			res = append(res, i)
		}
	}
	return res
}

func sortString(s string) string {
	sBytes := []byte(s)
	sort.Slice(sBytes, func(i, j int) bool {
		return sBytes[i] < sBytes[j]
	})
	return string(sBytes)
}

func main() {
	fmt.Println(findAnagramsLogN("cbaebabacd", "abc"))
	fmt.Println(findAnagramsN("cbaebabacd", "abc"))
}
