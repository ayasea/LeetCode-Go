package leetcode

import (
	"fmt"
	"testing"
)

type question392 struct {
	para392
	ans392
}

type para392 struct {
	one string
	ano string
}

type ans392 struct {
	one bool
}

func Test_Problem392(t *testing.T) {
	qs := []question392{
		{
			para392{"abc", "ahbgdc"},
			ans392{true},
		},
		{
			para392{"axc", "ahbgdc"},
			ans392{false},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 392-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans392, q.para392
		res := isSubsequenceV2(p.one, p.ano)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
