package codepractice

import (
	"fmt"
)

/*
70.爬楼梯
*/

// 递归
func climbStairs1(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

// 动态规划
func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func HandlePratice07Entry() {
	fmt.Println("n1: ", climbStairs(1))
	fmt.Println("n1: ", climbStairs(2))
	fmt.Println("n1: ", climbStairs(3))

}
