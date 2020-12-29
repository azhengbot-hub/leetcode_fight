package leetcode_greedy_test

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=330 lang=golang
 *
 * [330] 按要求补齐数组
 */

// @lc code=start
func minPatches(nums []int, n int) int {
	sumNums := 0
	index := 0
	count := 0

	for sumNums < n {
		if index < len(nums) && nums[index] <= sumNums+1 {
			sumNums += nums[index]
			index++
		} else {
			count++
			sumNums = sumNums + sumNums + 1
		}
	}
	return count
}

// @lc code=end

func TestFunc330(t *testing.T) {
	res := minPatches([]int{1, 5, 10}, 20)
	fmt.Println(res)
}
