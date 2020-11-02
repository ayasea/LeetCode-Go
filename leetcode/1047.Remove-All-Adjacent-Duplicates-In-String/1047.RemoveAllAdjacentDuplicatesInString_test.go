package leetcode

import (
	"fmt"
	"testing"
)

type question1047 struct {
	para1047
	ans1047
}

type para1047 struct {
	s string
}

type ans1047 struct {
	one string
}

func Test_Problem1047(t *testing.T) {
	qs := []question1047{
		{
			para1047{"abbaca"},
			ans1047{"ca"},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 1047-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans1047, q.para1047
		res := removeDuplicates(p.s)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
