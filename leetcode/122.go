// Package main 2024/1/22
package main

import (
	"fmt"
	"math"
)

/**
122. 买卖股票的最佳时机 II
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/
给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售
输入：prices = [7,1,5,3,6,4]
输出：7
解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3 。
     总利润为 4 + 3 = 7 。
**/
func maxProfit(prices []int) int {
	f0 := 0
	f1 := math.MinInt

	for _, p := range prices {
		f0 = max(f0, f1+p)
		f1 = max(f1, f0-p)
	}

	return f0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
