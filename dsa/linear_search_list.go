package main

func linearSearch(list []int, value int) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == value {
			return true
		}
	}
	
	return false
}