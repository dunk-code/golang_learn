package split

import (
	"strings"
)

func Split(s, sep string) (result []string) {
	// 优化
	result = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep)
	n := len(sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+n:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return result
}
