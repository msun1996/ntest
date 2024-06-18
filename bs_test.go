package ntest

/**
搜索插入位置
二分查找
*/

func searchInsert(nums []int, target int) int {
	pos := 0
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target <= nums[mid] {
			pos = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return pos
}
