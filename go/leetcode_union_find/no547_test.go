package leetcode_union_find_test

import (
	"fmt"
	"testing"
)

func findCircleNum(isConnected [][]int) int {
	father := make(map[int]int, len(isConnected))
	ans := 0

	add := func(x int) {
		if _, ok := father[x]; !ok {
			father[x] = -1
			ans++
		}
	}

	find := func(x int) int {
		root := x

		for father[root] != -1 {
			root = father[root]
		}

		for x != root {
			originalFather := father[x]
			father[x] = root
			x = originalFather
		}

		return root

	}

	merge := func(x, y int) {
		rootX, rootY := find(x), find(y)

		if rootX != rootY {
			father[rootX] = rootY
			ans--
		}
	}

	for i := 0; i < len(isConnected); i++ {
		add(i)

		for j := 0; j < i; j++ {
			if isConnected[i][j] == 1 {
				merge(i, j)
			}
		}
	}

	return ans
}

func TestFunc547(t *testing.T) {
	res := findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}})
	fmt.Println(res)
}
