package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question1441 struct {
	para1441
	ans1441
}

type para1441 struct {
	one []int
	n   int
}

type ans1441 struct {
	one []string
}

func Test_Problem1441(t *testing.T) {
	qs := []question1441{
		{
			para1441{[]int{1, 3}, 3},
			ans1441{[]string{"Push", "Push", "Pop", "Push"}},
		},
		{
			para1441{[]int{1, 2, 3}, 3},
			ans1441{[]string{"Push", "Push", "Push"}},
		},
		{
			para1441{[]int{1, 2}, 4},
			ans1441{[]string{"Push", "Push"}},
		},
		{
			para1441{[]int{2, 3, 4}, 4},
			ans1441{[]string{"Push", "Pop", "Push", "Push", "Push"}},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 1441-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans1441, q.para1441
		res := buildArray(p.one, p.n)
		if reflect.DeepEqual(ans.one, res) {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
