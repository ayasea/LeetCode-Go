package leetcode

import (
	"fmt"
	"testing"
)

type question242 struct {
	para242
	ans242
}

type para242 struct {
	one     string
	another string
}

type ans242 struct {
	one bool
}

func Test_Problem242(t *testing.T) {
	qs := []question242{
		{
			para242{"anagram", "nagaram"},
			ans242{true},
		},
		{
			para242{"rat", "car"},
			ans242{false},
		},
	}

	fmt.Printf("-----------------------------Leetcode Problem 242-----------------------------\n\n")

	for _, q := range qs {
		ans, p := q.ans242, q.para242
		res := isAnagram(p.one, p.another)
		if ans.one == res {
			fmt.Printf("[input]: %-6v  [output]: %-6v  PASS\n", p, res)
		} else {
			fmt.Printf("[input]: %-6v  [output]: %-6v  FAIL\n", p, res)
		}
	}
	fmt.Printf("\n")
}
