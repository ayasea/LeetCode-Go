package leetcode

import (
	"fmt"
	"testing"
)

type question874 struct {
	para874
	ans874
}

type para874 struct {
	one     []int
	another [][]int
}

type ans874 struct {
	one int
}

func Test_Problem874(t *testing.T) {
	qs := []question874{
		{
			para874{[]int{4, -1, 3}, [][]int{}},
			ans874{25},
		},
		{
			para874{[]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}},
			ans874{65},
		},
		{
			para874{[]int{-2, 8, 3, 7, -1}, [][]int{{-4, -1}, {1, -1}, {1, 4}, {5, 0}, {4, 5}, {-2, -1}, {2, -5}, {5, 1}, {-3, -1}, {5, -3}}},
			ans874{324},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 874-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans874, q.para874
		res := robotSim(p.one, p.another)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
