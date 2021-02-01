package leetcode_array_test

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	var sum int = 1
	var maxSum int

	if len(nums) <= 1 {
		return len(nums)
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			sum++
			continue
		}
		if sum > maxSum {
			maxSum = sum
		}

		sum = 1

	}

	if sum > maxSum {
		maxSum = sum
	}

	return maxSum
}

// @lc code=end

func TestFunc674(t *testing.T) {
	res := findLengthOfLCIS([]int{1, 3, 5, 4, 7})
	fmt.Println(res)
}
