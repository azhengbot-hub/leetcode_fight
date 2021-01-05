package leetcode_array_test

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	var targetList []int
	for _, vList := range matrix {
		if target >= vList[0] && target <= vList[len(vList)-1] {
			targetList = vList
			break
		}
	}
	for _, v := range targetList {
		if target == v {
			return true
		}
	}
	return false
}

// @lc code=end

func TestFun74(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	res := searchMatrix(matrix, target)
	fmt.Println(res)
}
