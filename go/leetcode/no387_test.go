package leetcode_test

import (
	"fmt"
	"testing"
)

func firstUniqChar(s string) int {
	resMap := make(map[int]int)
	for _, v := range s {
		if _, ok := resMap[int(v)]; ok {
			resMap[int(v)]++
		} else {
			resMap[int(v)] = 1
		}
	}
	for i, v := range s {
		if val, _ := resMap[int(v)]; val == 1 {
			return i
		}
	}
	return -1
}
func TestFun387(t *testing.T) {
	res := firstUniqChar("leetcode")
	fmt.Println(res)
}
