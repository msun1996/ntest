package ntest

import (
	"strconv"
)

/*
二进制求和
*/

func addBinary(a string, b string) string {
	sum := ""
	la := len(a)
	lb := len(b)
	lmax := getMax(la, lb)

	flag := 0
	for i := 0; i < lmax; i++ {
		ai, bi := 0, 0
		if la > i {
			ai, _ = strconv.Atoi(string(a[la-i-1]))
		}
		if lb > i {
			bi, _ = strconv.Atoi(string(b[lb-i-1]))
		}
		n := (ai + bi + flag) % 2
		sum = strconv.Itoa(n) + sum
		flag = (ai + bi + flag) / 2
	}
	if flag > 0 {
		sum = strconv.Itoa(flag) + sum
	}
	return sum
}
