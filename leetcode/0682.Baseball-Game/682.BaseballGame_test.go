package leetcode

import (
	"fmt"
	"testing"
)

type question682 struct {
	para682
	ans682
}

type para682 struct {
	one []string
}

type ans682 struct {
	one int
}

func Test_Problem682(t *testing.T) {
	qs := []question682{
		{
			para682{[]string{"5", "2", "C", "D", "+"}},
			ans682{30},
		},
		{
			para682{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}},
			ans682{27},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 682-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans682, q.para682
		res := calPoints(p.one)
		if ans.one == res {
			fmt.Printf("[input]: %-2v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-2v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
