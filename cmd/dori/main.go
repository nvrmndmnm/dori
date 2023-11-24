package main

import "fmt"

func main() {
	var r bool
	fmt.Println("hello Dori")

	r = linearSearch([]int{1, 2, 3}, 3)
	fmt.Println(r)

	r = binarySearchList([]int{1, 2, 3, 5, 9, 12, 13, 15}, 18)
	fmt.Println(r)

	i := twoCrystalBalls([]bool{false, false, false, false, true, true, true, true, true, true, true, true, true})
	fmt.Println(i)

	b:= bubbleSort([]int{6, 3, 7, 4, 2})
	fmt.Println(b)
}
