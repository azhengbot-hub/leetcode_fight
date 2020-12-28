package leetcode_stock_test

import (
	"fmt"
	"math"
	"testing"
)

/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit188(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	haveStock := make([][]int, n)
	noStock := make([][]int, n)

	for i := 0; i < n; i++ {
		haveStock[i] = make([]int, k+1)
		noStock[i] = make([]int, k+1)
	}
	haveStock[0][0] = -prices[0]
	noStock[0][0] = 0

	for i := 1; i < k; i++ {
		haveStock[0][i] = math.MinInt64 / 2
		noStock[0][i] = math.MinInt64 / 2
	}

	for i := 1; i < n; i++ {
		haveStock[i][0] = max(haveStock[i-1][0], noStock[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			haveStock[i][j] = max(haveStock[i-1][j], noStock[i-1][j]-prices[i])
			noStock[i][j] = max(noStock[i-1][j], haveStock[i-1][j-1]+prices[i])
		}
	}
	fmt.Println(haveStock)
	fmt.Println(noStock)
	return maxList(noStock[n-1]...)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxList(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > res {
			res = a[i]
		}
	}
	return res
}

// @lc code=end

func TestFunc188(t *testing.T) {
	res := maxProfit188(2, []int{2, 4, 1})
	fmt.Println(res)

}
