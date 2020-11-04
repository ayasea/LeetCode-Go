package leetcode

func minCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for _, c := range position {
		if c%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return min(odd, even)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
