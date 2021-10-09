package main

import (
	"unicode"
)

func Unpack(str string) (string, error) {
	var result string
	var escapeFlag bool
	var prevEscape bool
	for i, v := range str {
		if string(v) == "/" && !escapeFlag {
			escapeFlag = true
			continue
		}
		if escapeFlag {
			result += string(v)
		} else if unicode.IsDigit(v) {
			if i > 0 && (!unicode.IsDigit(rune(str[i-1])) || prevEscape) {
				for j := 0; j < int(v-'0')-1; j++ {
					result += string(str[i-1])
				}
			}
		} else {
			result += string(v)
		}
		if prevEscape {
			prevEscape = false
		}

		if escapeFlag {
			escapeFlag = false
			prevEscape = true
		}
	}
	return result, nil
}

func main() {
	test3 := "qwe/45ab"
	test4 := "qwe//5"
	Unpack(test3)
	Unpack(test4)
}
