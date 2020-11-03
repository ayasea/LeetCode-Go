package leetcode

import (
	"fmt"
	"testing"
)

type question122 struct {
	para122
	ans122
}

type para122 struct {
	one []int
}

type ans122 struct {
	one int
}

func Test_Problem122(t *testing.T) {
	qs := []question122{
		{
			para122{[]int{7, 1, 5, 3, 6, 4}},
			ans122{7},
		},
		{
			para122{[]int{1, 2, 3, 4, 5}},
			ans122{4},
		},
		{
			para122{[]int{7, 6, 4, 3, 1}},
			ans122{0},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 122-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans122, q.para122
		res := maxProfit(p.one)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
