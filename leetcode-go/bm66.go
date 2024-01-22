// Package main 2024/1/23
package main

import (
	"fmt"
)

/**
BM66 最长公共子串
https://www.nowcoder.com/share/jump/68216111705945764340
给定两个字符串str1和str2,输出两个字符串的最长公共子串
题目保证str1和str2的最长公共子串存在且唯一
**/
func lcs(str1 string, str2 string) string {
	n, m := len(str1), len(str2)
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	res, end := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if str1[i] == str2[j] {
				if i-1 < 0 || j-1 < 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			}
			if res < dp[i][j] {
				res = dp[i][j]
				end = i
			}
		}
	}
	if res == 0 {
		return ""
	}

	return str1[end-res+1 : end+1]
}

func main() {
	fmt.Println(lcs("1AB2345CD", "12345EF"))
}
