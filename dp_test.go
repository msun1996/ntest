package ntest

/*
*
最大子数组和
动态规划O(n)/O(1)
*/
func maxSubArray(nums []int) int {
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		// 滚动数组，num[i] 左侧加起来最大。
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}

	return maxNum
}
