package main

import (
	"testing"
)

func TestBinarySearchList(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	testCases := []struct {
		name     string
		search   int
		expected bool
	}{
		{"Element exists in array", 69, true},
		{"Element does not exist in array", 1336, false},
		{"Element exists in array", 69420, true},
		{"Element does not exist in array", 69421, false},
		{"Element exists in array", 1, true},
		{"Element does not exist in array", 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := binarySearchList(foo, tc.search)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}