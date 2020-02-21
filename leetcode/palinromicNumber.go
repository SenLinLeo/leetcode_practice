package main

import (
	"strings"
	"unicode"
)

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。
*/

func isPalindrome(s string) bool {

	tmp := strings.ToLower(s)
	var newStr []rune

	for i := 0; i < len(tmp); i ++ {
		if unicode.IsLetter(rune(tmp[i])) || unicode.IsNumber(rune(tmp[i])){
			newStr = append(newStr, rune(tmp[i]))
		}
	}

	i, j := 0, len(newStr) - 1
	for i < j {
		if newStr[i] != newStr[j] {
			return false
		}
		i ++
		j --
	}
	return true

}

func main(){
	flag := isPalindrome("ada ")
	println(flag)
	B := uint8(0)
	println(B)
}