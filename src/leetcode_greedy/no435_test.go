package leetcode_greedy_test

import (
	"fmt"
	"sort"
	"testing"
)

/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[i][1] })
	return 1
}

// @lc code=en

func TestFunc435(t testing.T) {
	res := eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})
	fmt.Println(res)
}
