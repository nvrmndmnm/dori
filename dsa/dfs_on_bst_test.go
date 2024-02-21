package main

import "testing"

func TestDFSonBST(t *testing.T) {
	result1 := DFS(tree, 45)
	expected1 := true
	if result1 != expected1 {
		t.Errorf("Test case 1 failed. Expected %t, got %t", expected1, result1)
	}

	result2 := DFS(tree, 7)
	expected2 := true
	if result2 != expected2 {
		t.Errorf("Test case 2 failed. Expected %t, got %t", expected2, result2)
	}

	result3 := DFS(tree, 69)
	expected3 := false
	if result3 != expected3 {
		t.Errorf("Test case 3 failed. Expected %t, got %t", expected3, result3)
	}
}