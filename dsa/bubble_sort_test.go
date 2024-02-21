package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}

	expected := []int{3, 4, 7, 9, 42, 69, 420}

	bubbleSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v but got %v", expected, arr)
	}
}
