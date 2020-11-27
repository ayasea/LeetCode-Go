package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question350 struct {
	para350
	ans350
}

type para350 struct {
	one     []int
	another []int
}

type ans350 struct {
	one []int
}

func Test_Problem350(t *testing.T) {
	qs := []question350{
		{
			para350{[]int{1, 2, 2, 1}, []int{2, 2}},
			ans350{[]int{2, 2}},
		},
		{
			para350{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
			ans350{[]int{4, 9}},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 350-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans350, q.para350
		res := intersectV2(p.one, p.another)
		if reflect.DeepEqual(ans.one, res) {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
