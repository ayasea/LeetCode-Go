package leetcode

import "sort"

func lastStoneWeight(stones []int) int {
	for len(stones) > 0 {
		length := len(stones)
		if length == 1 {
			return stones[0]
		}
		sort.Ints(stones)
		last := stones[length-1] - stones[length-2]
		stones = stones[:length-2]
		stones = append(stones, last)
	}
	return 0
}
