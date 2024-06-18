package ntest

/*
*
合并顺序两数组，合并至数组nums1
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 把nums2所有数据加到nums1，就算结束
	for n > 0 {
		if m-1 < 0 || nums2[n-1] >= nums1[m-1] {
			nums1[n+m-1] = nums2[n-1]
			n--
		} else {
			nums1[n+m-1] = nums1[m-1]
			m--
		}
	}
}

/*
*
移除元素
*/
func removeElement(nums []int, val int) int {
	left := 0
	for _, num := range nums {
		if num != val {
			nums[left] = num
			left++
		}
	}
	return left
}
