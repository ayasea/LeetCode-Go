package leetcode

func removeDuplicates(S string) string {
	if len(S) == 0 {
		return ""
	}
	stack := []byte{S[0]}
	for i := 1; i < len(S); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}
