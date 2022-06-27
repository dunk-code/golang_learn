package main

import (
	"fmt"
	"strings"
)

func lengthLongestPath(input string) int {
	idx := strings.Index(input, "\n")
	if idx == -1 {
		index := strings.Index(input, ".")
		if index != -1 {
			return len(input)
		} else {
			return 0
		}
	}
	maxDepth := 0
	root := input[:idx]
	mp := make(map[int]string)
	mp[0] = root
	for idx != -1 {
		startIdx := idx + 1
		depth := 0
		for startIdx+1 < len(input) && input[startIdx:startIdx+1] == fmt.Sprintf("\t") {
			depth++
			startIdx += 1
		}
		input = input[startIdx:]
		idx = strings.Index(input, "\n")
		prePath := mp[depth-1]
		if depth != 0 {
			prePath += "/"
		}
		if idx == -1 {
			prePath += input
		} else {
			prePath += input[:idx]
		}
		if strings.Index(prePath, ".") != -1 {
			maxDepth = max(maxDepth, depth)
		}
		mp[depth] = prePath
		idx = strings.Index(input, "\n")
	}
	ans := mp[maxDepth]
	if strings.Index(ans, ".") == -1 {
		return 0
	}
	return len(ans)
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	print(lengthLongestPath("file1.txt\nfile2.txt\nlongfile.txt"))
}
