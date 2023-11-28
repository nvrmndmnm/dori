package main

import (
	"testing"
)

func TestBFSGraphMatrix(t *testing.T) {
	// matrix2
	// {0, 3, 1, 0, 0, 0, 0}
	// {0, 0, 0, 0, 1, 0, 0}
	// {0, 0, 7, 0, 0, 0, 0}
	// {0, 0, 0, 0, 0, 0, 0}
	// {0, 1, 0, 5, 0, 2, 0}
	// {0, 0, 18, 0, 0, 0, 1}
	// {0, 0, 0, 1, 0, 0, 1}

	expectedPath := []int{0, 1, 4, 5, 6}
	result := BFSGraphMatrix(matrix2, 0, 6)

	if !intSlicesEqual(result, expectedPath) {
		t.Errorf("Expected path %v, but got %v", expectedPath, result)
	}

	if result := BFSGraphMatrix(matrix2, 6, 0); result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}
