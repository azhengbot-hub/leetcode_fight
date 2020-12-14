package leetcode_test

import (
	"fmt"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	// 给字符串排序
	// mp := map[string][]string{}
	// for _, str := range strs {
	// 	s := []byte(str)
	// 	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	// 	sortedStr := string(s)
	// 	mp[sortedStr] = append(mp[sortedStr], str)
	// }

	res := make(map[[26]int][]string)
	var resList [][]string

	for _, v := range strs {
		sum := [26]int{}
		for _, strInt := range v {
			sum[strInt-'a']++
		}
		if _, ok := res[sum]; ok {
			res[sum] = append(res[sum], v)
		} else {
			res[sum] = []string{v}
		}
	}
	fmt.Println(res)

	for _, v := range res {
		resList = append(resList, v)
	}
	return resList
}

func TestFunc49(t *testing.T) {
	res := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(res)
}
