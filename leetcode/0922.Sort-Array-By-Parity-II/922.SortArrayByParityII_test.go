package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question922 struct {
	para922
	ans922
}

type para922 struct {
	one []int
}

type ans922 struct {
	one []int
}

func Test_Problem922(t *testing.T) {
	qs := []question922{
		{
			para922{[]int{4, 2, 5, 7}},
			ans922{[]int{4, 5, 2, 7}},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 922-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans922, q.para922
		res := sortArrayByParityIIV2(p.one)
		if reflect.DeepEqual(ans.one, res) {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
