package leetcode

import (
	"fmt"
	"testing"
)

type question496 struct {
	para496
	ans496
}

type para496 struct {
	one     []int
	another []int
}

type ans496 struct {
	one []int
}

func Test_Problem496(t *testing.T) {
	qs := []question496{
		{
			para496{[]int{4, 1, 2}, []int{1, 3, 4, 2}},
			ans496{[]int{-1, 3, -1}},
		},
		{
			para496{[]int{2, 4}, []int{1, 2, 3, 4}},
			ans496{[]int{3, -1}},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 496-----------------------------\n")

	for _, q := range qs {
		_, p := q.ans496, q.para496
		fmt.Printf("[input] : %v              [output] : %v\n", p, nextGreaterElement(p.one, p.another))
	}
	fmt.Printf("\n\n\n")
}
