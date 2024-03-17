package leetcode

import "fmt"

func Run3Sum() {
	testcase := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(testcase))
}

func threeSum(nums []int) [][]int {

	numsMap := make(map[int]int)
	resMap := make(map[[3]int]int)
	for i := range nums {
		numsMap[nums[i]]++
	}

	for i := range nums {
		triplet := [3]int{10e5, 10e5, 10e5}
		key1 := nums[i]
		numsMap[key1]--
		for k,v := range numsMap {
			if v > 0 {
				triplet[1] = k
				numsMap[k]--
			}
		}
		triplet[0] = key1
		resMap[triplet]++
	}

	return [][]int{{}}
}
