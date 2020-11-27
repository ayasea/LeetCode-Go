package leetcode

import (
	"fmt"
	"testing"
)

type question976 struct {
	para976
	ans976
}

type para976 struct {
	one []int
}

type ans976 struct {
	one int
}

func Test_Problem976(t *testing.T) {
	qs := []question976{
		{
			para976{[]int{2, 1, 2}},
			ans976{5},
		},
		{
			para976{[]int{1, 2, 1}},
			ans976{0},
		},
		{
			para976{[]int{3, 2, 3, 4}},
			ans976{10},
		},
		{
			para976{[]int{3, 6, 2, 3}},
			ans976{8},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 976-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans976, q.para976
		res := largestPerimeter(p.one)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
