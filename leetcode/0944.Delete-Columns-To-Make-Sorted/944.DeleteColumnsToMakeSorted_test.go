package leetcode

import (
	"fmt"
	"testing"
)

type question944 struct {
	para944
	ans944
}

type para944 struct {
	one []string
}

type ans944 struct {
	one int
}

func Test_Problem944(t *testing.T) {
	qs := []question944{
		{
			para944{[]string{"cba", "daf", "ghi"}},
			ans944{1},
		},
		{
			para944{[]string{"a", "b"}},
			ans944{0},
		},
		{
			para944{[]string{"zyx", "wvu", "tsr"}},
			ans944{3},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 944-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans944, q.para944
		res := minDeletionSize(p.one)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
