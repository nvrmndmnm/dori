package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}

	QuickSort(arr)

	expected := []int{3, 4, 7, 9, 42, 69, 420}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Got %v, expected %v", arr, expected)
	}
}