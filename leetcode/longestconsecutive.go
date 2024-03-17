package leetcode

import (
	"fmt"
	"sort"
)

func RunLongestConsecutive() {
	// testcase := []int{100, 4, 200, 1, 3, 2}
	// testcase := []int{1, 2, 0, 1}
	// testcase := []int{}
	// testcase := []int{0}
	// testcase := []int{0, -1}
	// testcase := []int{0, 1, 2, 4, 8, 5, 6, 7, 9, 3, 55, 88, 77, 99, 999999999}
	// testcase := []int{12, 54, 15, 14, 16, 10, 11, 12, 13, 22, 23, 24, 25, 26}
	testcase := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(testcase))
}

func longestConsecutive(nums []int) int {
	count := 1
	longest := 1
	numsMap := make(map[int]int, len(nums))

	for i := range nums {
		numsMap[nums[i]] = nums[i]
	}
	uniqueNums := make([]int, 0)
	for _, num := range numsMap {
		uniqueNums = append(uniqueNums, num)
	}

	sort.Slice(uniqueNums, func(i, j int) bool {
		return uniqueNums[i] < uniqueNums[j]
	})

	for i := 0; i < len(uniqueNums)-1; i++ {
		if uniqueNums[i] == uniqueNums[i+1]-1 {
			count++
		} else {
			count = 1
		}

		if count > longest {
			longest = count
		}
	}

	if len(uniqueNums) == 0 {
		longest = 0
	}
	return longest
}
