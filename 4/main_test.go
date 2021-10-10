package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type addTest struct {
	input    string
	expected []string
}

func TestTop10(t *testing.T) {
	myList := NewLinkedList()
	myList.PushFront(10)
	assert.Equal(t, 1, myList.Len(), "PushFront, List len should be 1")
	assert.Equal(t, 10, myList.First().Value, "First should be 10")
	assert.Equal(t, 10, myList.Last().Value, "Last should be 10")
	myList.PushBack(20)
	assert.Equal(t, 2, myList.Len(), "PushBack, List len should be 2")
	assert.Equal(t, 10, myList.First().Value, "First should be 10")
	assert.Equal(t, 20, myList.Last().Value, "Last should be 20")
	myList.PushBack(30)
	myList.PushBack(40)
	myList.Remove(*myList.Last().Prev)
	assert.Equal(t, 3, myList.Len(), "Deleted element when len was 4, List len should be 3")
	assert.Equal(t, 10, myList.First().Value, "First should be 10")
	assert.Equal(t, 40, myList.Last().Value, "Last should be 20")

}
