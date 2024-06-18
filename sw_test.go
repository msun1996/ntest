package ntest

import "math"

/**
长度最小的子数组和>=目标值，子数组的长度
*/

func minSubArrLen(nums []int, target int) int {
	minLen := math.MaxInt

	start, end := 0, 0
	sum := 0

	// end 位移动至末尾
	for end < len(nums) {
		sum += nums[end]
		// sum 大于target 时，计算长度
		for sum >= target {
			minLen = getMinNum(minLen, end-start+1)
			// 移位start值
			sum -= nums[start]
			start++
		}
		end++
	}

	// 判断防止数组总和也没目标值大
	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}

func getMinNum(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}

/*
*
无重复字符的最长子串
*/
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	left, right := 0, 0
	maps := map[byte]int{}
	for right < len(s) {
		if pos, ok := maps[s[right]]; ok {
			for i := left; i <= pos; i++ {
				delete(maps, s[i])
			}
			left = pos + 1
		}
		maps[s[right]] = right
		maxLen = getMax(maxLen, right-left)
		right++
	}
	return maxLen
}
