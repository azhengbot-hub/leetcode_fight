package leetcode_array_test

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 */

// @lc code=start
func peakIndexInMountainArray(arr []int) int {
	var res int
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			continue
		} else {
			res = i - 1
			break
		}
	}
	return res
}

// @lc code=end

func TestFunc852(t *testing.T) {
	res := peakIndexInMountainArray([]int{3, 4, 5, 1})
	fmt.Println(res)
}
