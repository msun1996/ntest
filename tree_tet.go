package ntest

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
获取树的最大深度
深度搜索 O(n)/O(h)
*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getMax(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func getMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}

/*
*
二叉树右视图
前序遍历 中左右
层序遍历
*/
func rightSideView(root *TreeNode) []int {
	arrs := [][]int{}
	var dfs func(root *TreeNode, height int)
	dfs = func(root *TreeNode, height int) {
		if root == nil {
			return
		}
		if len(arrs) <= height {
			arrs = append(arrs, []int{root.Val})
		} else {
			arrs[height] = append(arrs[height], root.Val)
		}
		dfs(root.Left, height+1)
		dfs(root.Right, height+1)
	}
	dfs(root, 0)
	nums := []int{}
	for _, arr := range arrs {
		nums = append(nums, arr[len(arr)-1])
	}
	return nums
}

/*
*
前、中、后序遍历都是深度遍历 O(n)/O(h)

二叉搜索树的最小绝对差
二叉搜索树: 左叶子(及下所有节点) < 根叶子 < 右叶子
中序遍历 左中右 （大小依次增大）

	   4
	2     6

1   3
pre 1 -> 2 -> 3 -> 4 -> 6
root.Val 2 -> 3 -> 4 -> 6
*/
func getMinimumDifference(root *TreeNode) int {
	num := math.MaxInt
	pre := -100
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != -100 && root.Val-pre < num {
			num = root.Val - pre
		}
		pre = root.Val
		dfs(root.Right)
	}
	dfs(root)
	return num
}

/*
镜像树
*/
func isSymmetricTree(tree *TreeNode) bool {
	if tree == nil {
		return true
	}
	return isMirror(tree.Left, tree.Right)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return (left.Val == right.Val) && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}
