package ntest

/**
楼梯 1 或 2 层的上
1, 2, 3, 5.....
*/

func climbStairs(n int) int {
	f2, f1, f := 0, 0, 1
	for i := 1; i < n; i++ {
		f2 = f1
		f1 = f
		f = f1 + f2
	}
	return f
}
