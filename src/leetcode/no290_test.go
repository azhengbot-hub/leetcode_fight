package leetcode_test

import (
	"fmt"
	"strings"
	"testing"
)

func wordPattern(pattern string, s string) bool {
	resMap := map[int]string{}
	sList := strings.Split(s, " ")

	if len(pattern) != len(sList) {
		return false
	}
	for i, v := range sList {
		k := int(pattern[i])
		if resV, ok := resMap[k]; ok {
			if resV != v {
				return false
			}
		} else {
			resMap[k] = v
		}
	}
	resMapValue := map[string]int{}

	for k, v1 := range resMap {
		if v2, ok := resMapValue[v1]; ok {
			if v2 != k {
				return false
			}
		} else {
			resMapValue[v1] = k
		}
	}
	return true
}
func TestFunc290(t *testing.T) {
	res := wordPattern("aba", "cat cat cat dog")
	fmt.Println(res)
}
