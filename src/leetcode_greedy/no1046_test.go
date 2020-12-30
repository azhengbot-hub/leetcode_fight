package leetcode_greedy_test

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

/*
 * @lc app=leetcode.cn id=1046 lang=golang
 *
 * [1046] 最后一块石头的重量
 */

// @lc code=start
func lastStoneWeight(stones []int) int {
	for len(stones) >= 2 {
		sort.Sort(sort.Reverse(sort.IntSlice(stones)))
		leftStones := stones[2:]
		if stones[1] == stones[0] {
			stones = leftStones
		} else {
			stones = append(leftStones, int(math.Abs(float64(stones[0]-stones[1]))))
		}
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}

// @lc code=end

func TestFunc1046(t *testing.T) {
	res := lastStoneWeight([]int{2, 7, 4, 1, 8, 1})
	fmt.Println(res)
}
