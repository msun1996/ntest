package ntest

/*
数组中的第K个最大元素,使用堆排
*/

func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(nums []int, heapSiz int) {
	for i := heapSiz / 2; i >= 0; i-- {
		maxHeapify(nums, i, heapSiz)
	}
}

func maxHeapify(nums []int, i, heapSiz int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSiz && nums[l] > nums[largest] {
		largest = l
	}
	if r < heapSiz && nums[r] > nums[largest] {
		largest = r
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapify(nums, largest, heapSiz)
	}
}
