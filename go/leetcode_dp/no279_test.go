package leetcode_dp_test

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i, _ := range dp {
		dp[i] = n + 1
	}

	dp[0] = 0

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]

}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

func TestFun279(t *testing.T) {
	fmt.Println("start")
	res := numSquares(12)
	fmt.Println(res)
}
