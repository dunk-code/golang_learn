package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(sentence string) string {
	splits := strings.Split(sentence, " ")
	bytes := make([]byte, 0, len(splits))
	bytes = append(bytes, 'a')
	ans := make([]string, 0, len(splits))
	for _, str := range splits {
		var newStr string
		head := strings.ToLower(str)[0]
		if head == 'a' || head == 'e' || head == 'i' || head == 'o' || head == 'u' {
			newStr = str
		} else {
			newStr = str[1:] + str[:1]
		}
		newStr += "ma" + string(bytes)
		ans = append(ans, newStr)
		bytes = append(bytes, 'a')
	}
	return strings.Join(ans, " ")
}

func main() {
	latin := toGoatLatin("Each word consists of lowercase and uppercase letters only")
	fmt.Println(latin)
}
