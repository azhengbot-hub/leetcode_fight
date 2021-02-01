package leetcode_union_find_test

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=721 lang=golang
 *
 * [721] 账户合并
 */

// @lc code=start
func accountsMerge(accounts [][]string) [][]string {
	return [][]string{}
}

// @lc code=end

func TestFunc721(t *testing.T) {
	res := accountsMerge([][]string{{"John", "johnsmith@mail.com", "john00@mail.com"}, {"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"Mary", "mary@mail.com"}})
	fmt.Println(res)
}
