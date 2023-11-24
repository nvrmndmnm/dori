package main

import "testing"

func TestLinearSearch(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	tests := []struct {
		name     string
		search   int
		expected bool
	}{
		{"Found 69", 69, true},
		{"Not Found 1336", 1336, false},
		{"Found 69420", 69420, true},
		{"Not Found 69421", 69421, false},
		{"Found 1", 1, true},
		{"Not Found 0", 0, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := linearSearch(foo, test.search)
			if result != test.expected {
				t.Errorf("For %s, expected %v but got %v", test.name, test.expected, result)
			}
		})
	}
}