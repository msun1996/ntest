package ntest

import "math"

/*
*
三角形最小路径和
*/
func minimumTotal(triangle [][]int) int {
	for i := 0; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if i > 0 {
				if j-1 < 0 {
					triangle[i][j] += triangle[i-1][j]
				} else if j > len(triangle[i-1])-1 {
					triangle[i][j] += triangle[i-1][j-1]
				} else {
					triangle[i][j] += getMin(triangle[i-1][j-1], triangle[i-1][j])
				}
			}
		}
	}
	minNum := math.MaxInt
	for _, num := range triangle[len(triangle)-1] {
		minNum = getMin(minNum, num)
	}
	return minNum
}

func getMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}
