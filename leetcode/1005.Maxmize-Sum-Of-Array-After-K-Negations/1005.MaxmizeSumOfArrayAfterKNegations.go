package leetcode

import "sort"

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	sum := 0
	i := 0
	for i < len(A) && K > 0 {
		A[i] = -A[i]
		if A[i] > 0 && A[i] > A[i+1] {
			i++
		}
		K--
	}
	for _, v := range A {
		sum += v
	}
	return sum
}
