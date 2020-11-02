package leetcode

import (
	"fmt"
	"testing"
)

type question1598 struct {
	para1598
	ans1598
}

type para1598 struct {
	one []string
}

type ans1598 struct {
	one int
}

func Test_Problem1598(t *testing.T) {
	qs := []question1598{
		{
			para1598{[]string{"d1/", "d2/", "../", "d21/", "./"}},
			ans1598{2},
		},
		{
			para1598{[]string{"d1/", "d2/", "./", "d3/", "../", "d31/"}},
			ans1598{3},
		},
		{
			para1598{[]string{"d1/", "../", "../", "../"}},
			ans1598{0},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 1598-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans1598, q.para1598
		res := minOperations(p.one)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
