package main

import "fmt"

func findAndReplacePattern(words []string, pattern string) []string {
	n := len(pattern)
	ans := make([]string, 0, n)
	for _, word := range words {
		mp := map[uint8]uint8{}
		set := map[uint8]bool{}
		i := 0
		for ; i < n; i++ {
			if rw, ok := mp[word[i]]; ok && rw != pattern[i] || !ok && set[pattern[i]] {
				break
			}
			mp[word[i]] = pattern[i]
			set[pattern[i]] = true
		}
		if i == n {
			ans = append(ans, word)
		}
	}
	return ans
}

func main() {
	ans := findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb")
	for _, a := range ans {
		fmt.Println(a)
	}
}
