package leetcode_test

import (
	"fmt"
	"testing"
)

func predictPartyVictory(senate string) string {
	var radiant, dire []int

	for i, v := range senate {
		if v == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0] + len(senate))
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}
	if len(dire) > 0 {
		return "D"
	}

	return "R"
}
func TestFunc649(t *testing.T) {
	res := predictPartyVictory("RDD")
	fmt.Println(res)
}
