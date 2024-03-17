package leetcode

import "fmt"

func RunProductArray() {
	testcase := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(testcase))
}

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	for i := range nums {
		answer[i] = 1
	}
	
	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	tail := nums[len(nums)-1]
	for i := len(nums) - 2; i > -1; i-- {
		answer[i] *= tail
		tail *= nums[i]
	}

	return answer
}