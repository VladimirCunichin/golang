package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type addTest struct {
	input, expected string
}

var addTests = []addTest{
	{"a4bc2d5e", "aaaabccddddde"},
	{"abcd", "abcd"},
	{"45", ""},
	{"", ""},
	{"qwe/4/5", "qwe45"},
	{"qwe/45", "qwe44444"},
	{"qwe//5", "qwe/////"},
}

func TestUnpack(t *testing.T) {
	for _, test := range addTests {
		output, err := Unpack(test.input)
		if err != nil {
			t.Fatalf("error is not nil")
		}
		assert.Equal(t, test.expected, output, "Unpacking "+test.input+" expected "+test.expected)
	}
}
