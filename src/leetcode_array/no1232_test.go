package leetcode_array_test

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=1232 lang=golang
 *
 * [1232] 缀点成线
 */

// @lc code=start
func checkStraightLine(coordinates [][]int) bool {
	n := len(coordinates)

	if n <= 2 {
		return true
	}

	for i := 0; i+2 < n; i++ {
		r1 := float64(coordinates[i+2][0]-coordinates[i+1][0]) / float64(coordinates[i+2][1]-coordinates[i+1][1])
		r2 := float64(coordinates[i+1][0]-coordinates[i][0]) / float64(coordinates[i+1][1]-coordinates[i][1])
		fmt.Println(r1)
		fmt.Println(r2)
		if r1 == r2 {
			continue
		} else if coordinates[i+2][1] == coordinates[i+1][1] && coordinates[i+1][1] == coordinates[i][1] {
			continue
		} else {
			return false
		}
	}
	return true
}

// @lc code=end

func TestFunc1232(t *testing.T) {
	res := checkStraightLine([][]int{{-3, -2}, {-1, -2}, {2, -2}, {-2, -2}, {0, -2}})
	fmt.Println(res)
}
