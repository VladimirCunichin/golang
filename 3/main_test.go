package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type addTest struct {
	input    string
	expected []string
}

var addTests = []addTest{
	{"Hello hello hello, Hello, zero zero Zero one one, Test, Two Test Two, four test Four, six Six Six, hey hey hey hey nine nine nine ten ten ten ten wrong", []string{"hello", "hey", "ten", "test", "six", "nine", "zero", "one", "two", "four"}},
}

func TestTop10(t *testing.T) {
	for _, test := range addTests {
		assert.ElementsMatch(t, test.expected, Top10(test.input), "Unpacking "+test.input)
	}
}
