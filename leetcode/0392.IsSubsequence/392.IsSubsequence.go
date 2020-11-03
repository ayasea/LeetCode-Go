package leetcode

func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == n
}

func isSubsequenceV2(s string, t string) bool {
	var temp = false

	if s == "" {
		return true
	}
	for j := 0; j < len(t); j++ {
		if s[0] == t[j] {
			if len(s) == 1 {
				temp = true
				break
			}
			return isSubsequenceV2(s[1:len(s)], t[j+1:len(t)])
		}
	}
	if temp {
		return true
	}
	return false
}
