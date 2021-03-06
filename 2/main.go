package main

import (
	"errors"
	"fmt"
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
			} else {
				return "", errors.New("two digits in a row - incorrect input")
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
	test3 := "45"
	test4 := "qwe//5"
	res, err := Unpack(test3)
	if err != nil {
		fmt.Println(res, err)
	}
	fmt.Println(Unpack(test4))
}
