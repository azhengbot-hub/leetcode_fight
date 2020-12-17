package leetcode_test

import (
	"fmt"
	"math"
	"testing"
)

func maxProfit(prices []int, fee int) int {
	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = []int{0, 0}
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]+prices[i]-fee)))
		dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(dp[i-1][0]-prices[i])))
	}
	return dp[len(prices)-1][0]
}
func TestFunc714(t *testing.T) {
	res := maxProfit([]int{1, 3, 2, 8, 4, 9}, 2)
	fmt.Println(res)
}
