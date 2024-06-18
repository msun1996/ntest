package ntest

import (
	"sort"
	"strconv"
)

/**
区间汇总
["0", "1", "2", "4", "5", "7"] -> ["0->2", "4->5", "7"]
*/

func summaryRanges(nums []int) []string {
	strs := []string{}
	i := 0
	n := len(nums)
	for i < n {
		left := i
		for i < n && nums[i-1]+1 == nums[i] {
			i++
		}
		// 右边界需要退1
		right := i - 1
		str := strconv.Itoa(nums[left])
		if left < right {
			str = str + "->" + strconv.Itoa(nums[right])
		}
		strs = append(strs, str)
	}
	return strs
}

/*
*
合并区间
*/
func mergePart(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	arrs := [][]int{}
	arrs = append(arrs, intervals[0])
	for _, interval := range intervals[1:] {
		if interval[0] > arrs[len(arrs)-1][1] {
			arrs = append(arrs, interval)
		} else {
			left := arrs[len(arrs)-1][0]
			right := getMax(arrs[len(arrs)-1][1], interval[1])
			arrs[len(arrs)-1] = []int{left, right}
		}
	}
	return arrs
}
