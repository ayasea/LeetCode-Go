package leetcode

func minOperations(log []string) int {
	stack := []string{}
	for i := range log {
		if log[i] == "../" {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if log[i] != "./" {
			stack = append(stack, log[i])
		}
	}
	return len(stack)
}
