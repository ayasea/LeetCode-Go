package leetcode

import (
	"fmt"
	"testing"
)

type question1217 struct {
	para1217
	ans1217
}

type para1217 struct {
	one []int
}

type ans1217 struct {
	one int
}

func Test_Problem1217(t *testing.T) {
	qs := []question1217{
		{
			para1217{[]int{1, 2, 3}},
			ans1217{1},
		},
		{
			para1217{[]int{2, 2, 2, 3, 3}},
			ans1217{2},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 1217-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans1217, q.para1217
		res := minCostToMoveChips(p.one)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
