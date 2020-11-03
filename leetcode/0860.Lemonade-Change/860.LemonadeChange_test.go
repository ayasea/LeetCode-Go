package leetcode

import (
	"fmt"
	"testing"
)

type question860 struct {
	para860
	ans860
}

type para860 struct {
	one []int
}

type ans860 struct {
	one bool
}

func Test_Problem860(t *testing.T) {
	qs := []question860{
		{
			para860{[]int{5, 5, 5, 10, 20}},
			ans860{true},
		},
		{
			para860{[]int{5, 5, 10}},
			ans860{true},
		},
		{
			para860{[]int{10, 10}},
			ans860{false},
		},
		{
			para860{[]int{5, 5, 10, 10, 20}},
			ans860{false},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 860-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans860, q.para860
		res := lemonadeChange(p.one)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
