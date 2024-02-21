package main

import "testing"

func TestBFS(t *testing.T) {
	result1 := BFS(tree, 45)
	expected1 := true
	if result1 != expected1 {
		t.Errorf("Test case 1 failed. Expected %t, got %t", expected1, result1)
	}

	result2 := BFS(tree, 7)
	expected2 := true
	if result2 != expected2 {
		t.Errorf("Test case 2 failed. Expected %t, got %t", expected2, result2)
	}

	result3 := BFS(tree, 69)
	expected3 := false
	if result3 != expected3 {
		t.Errorf("Test case 3 failed. Expected %t, got %t", expected3, result3)
	}

	result4 := BFS2(tree, 45)
	expected4 := true
	if result4 != expected4 {
		t.Errorf("Test case 4 failed. Expected %t, got %t", expected4, result4)
	}

	result5 := BFS2(tree, 7)
	expected5 := true
	if result5 != expected5 {
		t.Errorf("Test case 5 failed. Expected %t, got %t", expected5, result5)
	}

	result6 := BFS2(tree, 69)
	expected6 := false
	if result6 != expected6 {
		t.Errorf("Test case 6 failed. Expected %t, got %t", expected6, result6)
	}
}