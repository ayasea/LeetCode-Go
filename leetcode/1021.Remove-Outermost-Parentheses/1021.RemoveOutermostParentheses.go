package leetcode

import "strings"

//解法一
func removeOuterParentheses(S string) string {
	now, current, ans := 0, "", ""
	for _, char := range S {
		if string(char) == "(" {
			now++
		} else if string(char) == ")" {
			now--
		}
		current += string(char)
		if now == 0 {
			ans += current[1 : len(current)-1]
			current = ""
		}
	}
	return ans
}

//解法二
func removeOuterParenthesesV2(S string) string {
	if len(S) == 0 {
		return ""
	}
	stack := []byte{S[0]}
	start := 1
	var res strings.Builder
	for i := 1; i < len(S); i++ {
		if len(stack) > 0 && S[i] == ')' && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				res.WriteString(S[start:i])
				start = i + 2
			}
		} else {
			stack = append(stack, S[i])
		}
	}
	return res.String()
}
