package leetcode

import (
	"fmt"
	"testing"
)

type question1544 struct {
	para1544
	ans1544
}

type para1544 struct {
	one string
}

type ans1544 struct {
	one string
}

func Test_Problem1544(t *testing.T) {
	qs := []question1544{
		{
			para1544{"leEeetcode"},
			ans1544{"leetcode"},
		},
		{
			para1544{"abBAcC"},
			ans1544{""},
		},
		{
			para1544{"s"},
			ans1544{"s"},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 1544-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans1544, q.para1544
		res := makeGood(p.one)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
