package leetcode_array_test

import (
	"fmt"
	"sort"
	"testing"
)

/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 */

// @lc code=start
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	var maxRes int

	start := nums[0] * nums[1] * nums[2]
	end := nums[n-1] * nums[n-2] * nums[n-3]
	other := nums[0] * nums[1] * nums[n-1]

	valueList := []int{start, end, other}

	for i, v := range valueList {
		if i == 0 {
			maxRes = v
		} else if v > maxRes {
			maxRes = v
		}
	}
	return maxRes

}

// @lc code=end

func TestFunc628(t *testing.T) {
	res := maximumProduct([]int{1, 2, 3, 4})
	fmt.Println(res)
}
