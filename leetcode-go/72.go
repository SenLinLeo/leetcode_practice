package main

/**
72. 编辑距离
https://leetcode.cn/problems/edit-distance/description/?envType=study-plan-v2&envId=top-100-liked
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
你可以对一个单词进行如下三种操作：
插入一个字符
删除一个字符
替换一个字符

示例 1：
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
**/
func minDistance(s string, t string) int {
	m := len(t)
	f := make([]int, m+1)
	for j := 1; j <= m; j++ {
		f[j] = j
	}
	for _, x := range s {
		pre := f[0]
		f[0]++
		for j, y := range t {
			if x == y {
				f[j+1], pre = pre, f[j+1]
			} else {
				f[j+1], pre = min(min(f[j+1], f[j]), pre)+1, f[j+1]
			}
		}
	}
	return f[m]
}

func main() {

}
