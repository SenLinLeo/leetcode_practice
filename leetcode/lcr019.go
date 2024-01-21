package main

/*
LCR 019. 验证回文串
给定一个非空字符串 s，请判断如果 最多从字符串中删除一个字符能否得到一个回文字符串。
https://leetcode.cn/problems/RQku0D/description/
*/
func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s[left:right]) || isPalindrome(s[left+1:right+1])
		}
		left++
		right--
	}
	return true
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r && s[l] == s[r] {
		l++
		r--
	}
	return l >= r
}

func main() {
	println(isPalindrome("ada"))
	println(isPalindrome("adac "))

}
