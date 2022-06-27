package _22__最长特殊序列_II

func findLUSlength(strs []string) int {
	ans := -1
next:
	for i, s := range strs {
		for j, t := range strs {
			if i != j && isSubseq(s, t) {
				continue next
			}
		}
		if len(s) > ans {
			ans = len(s)
		}
	}
	return ans
}

func isSubseq(s, t string) bool {
	ptS := 0
	for ptT := range t {
		if s[ptS] == t[ptT] {
			ptS++
			if ptS == len(s) {
				return true
			}
		}
	}
	return false
}
