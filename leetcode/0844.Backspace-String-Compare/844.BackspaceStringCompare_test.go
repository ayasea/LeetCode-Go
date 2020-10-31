package leetcode

import (
	"fmt"
	"testing"
)

type question844 struct {
	para844
	ans844
}

type para844 struct {
	s string
	t string
}

type ans844 struct {
	one bool
}

func Test_Problem844(t *testing.T) {
	qs := []question844{
		{
			para844{"ab#c", "ad#c"},
			ans844{true},
		},
		{
			para844{"ab##", "c#d#"},
			ans844{true},
		},
		{
			para844{"a##c", "#a#c"},
			ans844{true},
		},
		{
			para844{"a#c", "b"},
			ans844{false},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 844-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans844, q.para844
		res := backspaceCompare(p.s, p.t)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
