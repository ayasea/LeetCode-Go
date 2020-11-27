package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question349 struct {
	para349
	ans349
}

type para349 struct {
	one     []int
	another []int
}

type ans349 struct {
	one []int
}

func Test_problem349(t *testing.T) {
	qs := []question349{
		{

			para349{[]int{1, 2, 2, 1}, []int{2, 2}},
			ans349{[]int{2}},
		},
		{
			para349{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
			ans349{[]int{9, 4}},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 349-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans349, q.para349
		res := intersectionV2(p.one, p.another)
		if reflect.DeepEqual(ans.one, res) {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
