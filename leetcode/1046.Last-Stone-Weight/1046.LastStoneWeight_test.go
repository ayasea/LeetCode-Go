package leetcode

import (
	"fmt"
	"testing"
)

type question1046 struct {
	para1046
	ans1046
}

type para1046 struct {
	one []int
}

type ans1046 struct {
	one int
}

func Test_Problem1046(t *testing.T) {
	qs := []question1046{
		{
			para1046{[]int{2, 7, 4, 1, 8, 1}},
			ans1046{1},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 1046-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans1046, q.para1046
		res := lastStoneWeight(p.one)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
