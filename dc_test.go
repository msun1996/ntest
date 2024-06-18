package ntest

/**
分治
将有序数组转换为二叉搜索树。

*/

func sortedArrayToBST(nums []int) *TreeNode {
	var dfs func(nums []int, left, right int) *TreeNode
	dfs = func(nums []int, left, right int) *TreeNode {
		// 终止条件
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		// 中间值为根节点
		node := &TreeNode{Val: nums[mid]}
		// 左边数组为左节点
		node.Left = dfs(nums, left, mid-1)
		// 右边数组为右节点
		node.Right = dfs(nums, mid+1, right)
		return node
	}
	node := dfs(nums, 0, len(nums)-1)
	return node
}
