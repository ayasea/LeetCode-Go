package leetcode

import (
	"fmt"
	"testing"
)

type question1021 struct {
	para1021
	ans1021
}

type para1021 struct {
	one string
}

type ans1021 struct {
	one string
}

func Test_Problem1021(t *testing.T) {
	qs := []question1021{
		{
			para1021{"(()())(())"},
			ans1021{"()()()"},
		},
		{
			para1021{"(()())(())(()(()))"},
			ans1021{"()()()()(())"},
		},
		{
			para1021{"()()"},
			ans1021{""},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 1021-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans1021, q.para1021
		res := removeOuterParenthesesV2(p.one)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
