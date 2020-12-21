package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func minCostClimbingStairs(cost []int) int {
	resList := make([][2]int, len(cost))

	for i, v := range cost {
		if i == 0 {
			resList[i][0] = v
			resList[i][1] = v
		} else if i == 1 {
			resList[i][0] = cost[0] + v
			resList[i][1] = v
		} else {
			resList[i][0] = int(math.Min(float64(resList[i-1][0]), float64(resList[i-1][1]))) + v
			resList[i][1] = int(math.Min(float64(resList[i-2][0]), float64(resList[i-2][1]))) + v
		}
	}
	fmt.Println(resList)
	return int(
		math.Min(
			math.Min(float64(resList[len(cost)-1][0]), float64(resList[len(cost)-1][1])),
			math.Min(float64(resList[len(cost)-2][0]), float64(resList[len(cost)-2][1])),
		),
	)
}

func TestFunc746(t *testing.T) {
	res := minCostClimbingStairs([]int{1, 0, 0, 0})
	fmt.Println(res)
}
