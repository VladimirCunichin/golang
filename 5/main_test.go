package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type addTest struct {
	input    []func() error
	N        int
	M        int
	expected error
}

var addTests = []addTest{
	{[]func() error{f1, f1, f1, f1, f1, f1, f1, f1, f1, f1}, 4, 1, nil},
	{[]func() error{f1, f1, f1, f1, f1, ef1, ef1, ef1, ef1, ef1}, 4, 3, fmt.Errorf("too many errors, completed tasks: 4")},
	{[]func() error{f1, f1, f1, f1, f1, ef1, ef1, ef1, ef1, ef1}, 4, 6, nil},
	{[]func() error{ef1, f1, f1, f1, f1, f1, f1, f1, f1, f1}, 4, 1, fmt.Errorf("too many errors, completed tasks: 10")},
}

func TestRun(t *testing.T) {
	for _, test := range addTests {
		assert.Equal(t, Run(test.input, test.N, test.M), test.expected)
	}
}
