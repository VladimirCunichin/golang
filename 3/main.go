package main

import (
	"fmt"
	"sort"
	"strings"
)

//Top10 most repeated words in a string
func Top10(input string) []string {
	wordsMap := make(map[string]int)
	words := strings.Split(input, " ")
	for _, v := range words {
		if v == "-" {
			continue
		}
		current := v
		current = strings.TrimSuffix(current, ",")
		current = strings.ToLower(current)
		value, ok := wordsMap[current]
		if ok {
			wordsMap[current] = value + 1
		} else {
			wordsMap[current] = 1
		}
	}

	return rankMapStringInt(wordsMap)
}

func rankMapStringInt(values map[string]int) []string {
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range values {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	ranked := make([]string, len(values))
	for i, kv := range ss {
		ranked[i] = kv.Key
	}
	return ranked[:10]
}

func main() {
	test := "Hello, today i would like to talk about writing a hello world program hello, world Hello,"
	fmt.Println(Top10(test))
}
