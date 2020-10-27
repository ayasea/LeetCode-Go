package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem255(t *testing.T) {
	obj := Constructor225()
	fmt.Printf("obj = %v\n", obj)
	param1 := obj.Empty()
	fmt.Printf("param5 = %v\n", param1)
	obj.Push(2)
	fmt.Printf("obj = %v\n", obj)
	obj.Push(10)
	fmt.Printf("obj = %v\n", obj)
	param2 := obj.Pop()
	fmt.Printf("param2 = %v\n", param2)
	param3 := obj.Top()
	fmt.Printf("param3 = %v\n", param3)
	param4 := obj.Empty()
	fmt.Printf("param4 = %v\n", param4)
}
