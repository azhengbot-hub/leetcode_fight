package leetcode_greedy_test

import (
	"fmt"
	"math"
	"testing"
)

func candy(ratings []int) int {
	ratingsLength := len(ratings)
	left := make([]int, ratingsLength)
	right := make([]int, ratingsLength)
	ans := 0

	for i, _ := range left {
		left[i] = 1
	}
	for i := 1; i < ratingsLength; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}

	for i, _ := range right {
		right[i] = 1
	}
	for i := ratingsLength - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			right[i-1] = right[i] + 1
		}
	}

	for i := 0; i < ratingsLength; i++ {
		ans += int(math.Max(float64(left[i]), float64(right[i])))
	}

	return ans
}
func TestFunc135(t *testing.T) {
	res := candy([]int{1, 2, 87, 87, 87, 2, 1})
	fmt.Println(res)
}
