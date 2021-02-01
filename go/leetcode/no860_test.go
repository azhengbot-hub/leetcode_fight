package leetcode_test

import (
	"fmt"
	"testing"
)

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0

	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five < 1 {
				return false
			}
			five--
			ten++
		} else {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true

}
func TestFunc860(t *testing.T) {
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20}))
}
