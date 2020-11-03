package leetcode

import (
	"fmt"
	"testing"
)

type question455 struct {
	para455
	ans455
}

type para455 struct {
	one     []int
	another []int
}

type ans455 struct {
	one int
}

func TestProblem455(t *testing.T) {
	qs := []question455{
		{
			para455{[]int{1, 2, 3}, []int{1, 1}},
			ans455{1},
		},
		{
			para455{[]int{1, 2}, []int{1, 2, 3}},
			ans455{2},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 455-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans455, q.para455
		res := findContentChildren(p.one, p.another)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
