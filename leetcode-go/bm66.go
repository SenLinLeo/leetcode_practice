// Package main 2024/1/23
package main

func lcs(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {

			}
			dp[i][j] = dp[i-1][j-1] + 1
		}
	}

}

func main() {

}
