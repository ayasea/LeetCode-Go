package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question1030 struct {
	para1030
	ans1030
}

type para1030 struct {
	R  int
	C  int
	r0 int
	c0 int
}

type ans1030 struct {
	one [][]int
}

func Test_Problem1030(t *testing.T) {
	qs := []question1030{
		{
			para1030{1, 2, 0, 0},
			ans1030{[][]int{{0, 0}, {0, 1}}},
		},
		{
			para1030{2, 2, 0, 1},
			ans1030{[][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}}},
		},
		{
			para1030{2, 3, 1, 2},
			ans1030{[][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}}},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 1030-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans1030, q.para1030
		res := allCellsDistOrder(p.R, p.C, p.r0, p.c0)
		if reflect.DeepEqual(ans.one, res) {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
