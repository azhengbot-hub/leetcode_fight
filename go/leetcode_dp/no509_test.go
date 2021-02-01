package leetcode_dp_test

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		if i == 0 {
			dp[i] = 0
		} else if i == 1 {
			dp[i] = 1
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[n]
}

// @lc code=end

func TestFun509(t *testing.T) {
	res := fib(0)
	fmt.Println(res)
}
