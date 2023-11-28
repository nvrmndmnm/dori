package main

import (
	"testing"
)

func TestDFSGraphList(t *testing.T) {
	// list
	// {{1, 3}}
	// {{4, 1}}
	// {{3, 7}}
	// {}
	// {{1, 1}, {3, 5}, {5, 2}}
	// {{2, 18}, {6, 1}}
	// {{3, 1}}

	expected1 := []int{0, 1, 4, 5, 6}
	result := DFSGraphList(list2, 0, 6)

	if !intSlicesEqual(result, expected1) {
		t.Errorf("Expected path %v, but got %v", expected1, result)
	}

	if result := DFSGraphList(list2, 6, 0); result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}
