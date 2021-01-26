package leetcode_stack_test

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=1128 lang=golang
 *
 * [1128] 等价多米诺骨牌对的数量
 */

// @lc code=start
func numEquivDominoPairs(dominoes [][]int) int {
	res := [100]int{}
	ans := 0

	for _, v := range dominoes {
		if v[0] > v[1] {
			v[0], v[1] = v[1], v[0]
		}
		one := 10*v[0] + v[1]

		ans += res[one]

		res[one]++
	}
	fmt.Println(res)
	return ans
}

// @lc code=end

func TestFunc1128(t *testing.T) {
	res := numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}})
	fmt.Println(res)
}
