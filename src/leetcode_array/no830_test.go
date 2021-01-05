package leetcode_array_test

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=830 lang=golang
 *
 * [830] 较大分组的位置
 */

// @lc code=start
func largeGroupPositions(s string) [][]int {
	cnt := 1
	var res [][]int
	for i := range s {
		if i == len(s)-1 || s[i] != s[i+1] {
			if cnt >= 3 {
				res = append(res, []int{i - cnt + 1, i})
			}
			cnt = 1
		} else {
			cnt++
		}
	}
	return res
}

// @lc code=end

func TestFun830(t *testing.T) {
	res := largeGroupPositions("abcdddeeeeeeeaabbbcd")
	fmt.Println(res)
}
