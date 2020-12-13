package leetcode_test

import (
	"testing"
)

func uniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, m)
	}

	for i, vList := range dp {
		for j, _ := range vList {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}

		}
	}
	return dp[n-1][m-1]
}

func TestFunc(t *testing.T) {
	uniquePaths(3, 7)
}
