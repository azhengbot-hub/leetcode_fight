package leetcode_array_test

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心索引
 */

// @lc code=start
func pivotIndex(nums []int) int {
	allSum := sum(nums)

	for i, v := range nums {
		front := sum(nums[:i])
		after := allSum - front - v

		if front == after {
			return i
		}
	}
	return -1
}

func sum(nums []int) int {
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}

// @lc code=end

func TestFunc724(t *testing.T) {
	res := pivotIndex([]int{1, 7, 3, 6, 5, 6})
	fmt.Println(res)
}
