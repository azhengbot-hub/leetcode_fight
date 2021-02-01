package leetcode_array_test

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	if k > len(nums) {
		k = k % len(nums)
	}

	newNums := append(nums[len(nums)-k:], nums[:len(nums)-k]...)

	for i := 0; i < len(nums); i++ {
		nums[i] = newNums[i]
	}
}

// @lc code=end

func TestFunc189(t *testing.T) {
	rotate([]int{-1}, 2)
}
