package main

func binarySearchList(list []int, value int) bool {
	min := 0
	max := len(list)
	for min < max {
		m := min + (max - min) / 2
		v := list[m]
		if v == value {
			return true
		} else if v > value {
			max = m
		} else {
			min = m + 1
		}
	}
	return false
}
