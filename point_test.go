package ntest

import (
	"strings"
)

/*
*
回文字符串 "Abasa,asa,asaA1,sa112s" 忽略大小写和特殊字符
*/

func isLetterOrNum(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	left := 0
	right := len(str) - 1
	for left < right {
		if !isLetterOrNum(str[left]) {
			left++
			continue
		}
		if !isLetterOrNum(str[right]) {
			right--
			continue
		}
		if str[left] == str[right] {
			left++
			right--
			continue
		}
		return false
	}
	return true
}

/*
*
判断子序列
*/
func isSubsequence(s string, t string) bool {
	sLeft := 0
	for i := 0; i < len(s); i++ {
		if sLeft > len(t)-1 {
			return false
		}
		for s[i] != t[sLeft] {
			sLeft++
			if sLeft > len(t)-1 {
				return false
			}
		}
		sLeft++
	}
	return true
}
