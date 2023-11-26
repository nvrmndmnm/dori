package main

import "testing"

func TestCompareTrees(t *testing.T) {
	result1 := Compare(tree, tree)
	expected1 := true
	if result1 != expected1 {
		t.Errorf("Test case 1 failed. Expected %t, got %t", expected1, result1)
	}

	result2 := Compare(tree, tree2)
	expected2 := false
	if result2 != expected2 {
		t.Errorf("Test case 2 failed. Expected %t, got %t", expected2, result2)
	}
}