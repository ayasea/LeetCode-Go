package leetcode

import (
	"fmt"
	"testing"
)

type question1005 struct {
	para1005
	ans1005
}

type para1005 struct {
	one     []int
	another int
}

type ans1005 struct {
	one int
}

func Test_problem1005(t *testing.T) {
	qs := []question1005{
		{
			para1005{[]int{4, 2, 3}, 1},
			ans1005{5},
		},
		{
			para1005{[]int{3, -1, 0, 2}, 3},
			ans1005{6},
		},
		{
			para1005{[]int{2, -3, -1, 5, -4}, 2},
			ans1005{13},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 1005-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans1005, q.para1005
		res := largestSumAfterKNegations(p.one, p.another)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
