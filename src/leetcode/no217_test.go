package leetcode_test

import (
	"fmt"
	"testing"
)

func containsDuplicate(nums []int) bool {
	resMap := map[int]bool{}

	for _, v := range nums {
		if _, ok := resMap[v]; ok {
			return true
		} else {
			resMap[v] = true
		}
	}
	return false
}

func TestFunc217(t *testing.T) {
	res := containsDuplicate([]int{1, 2, 3, 1})
	fmt.Println(res)
}
