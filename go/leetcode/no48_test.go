package leetcode_test

import (
	"fmt"
	"testing"
)

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 1 + i; j < len(matrix[0]); j++ {
			fmt.Println(matrix[i][j], matrix[j][i])
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	fmt.Println(matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0])/2; j++ {
			matrix[i][j], matrix[i][len(matrix)-j-1] = matrix[i][len(matrix)-j-1], matrix[i][j]
		}
	}

	fmt.Println(matrix)
}
func TestFunc48(t *testing.T) {
	rotate([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}})
}
