package main

func minDeletionSize(strs []string) int {
	if strs == nil || len(strs) == 0 {
		return -1
	}
	bytes := make([]byte, 0, len(strs[0]))
	for i := 0; i < len(strs[0]); i++ {
		bytes = append(bytes, strs[0][i])
	}
	mp := map[int]bool{}
	ans := 0
	for i := 1; i < len(strs); i++ {
		str := strs[i]
		for j := 0; j < len(str); j++ {
			if !mp[j] {
				if str[j] < bytes[j] {
					mp[j] = true
					ans++
				}
				bytes[j] = str[j]
			}
		}
	}
	return ans
}

func main() {
	strings := []string{"rrjk", "furt", "guzm"}
	print(minDeletionSize(strings))
}
