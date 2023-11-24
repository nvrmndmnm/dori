package main

func bubbleSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		for j := 0; j < i; j++ { 
			if list[i] < list[j] {
				tmp := list[i]
				list[i] = list[j]
				list[j] = tmp
			}
		}
	}
	return list
}
