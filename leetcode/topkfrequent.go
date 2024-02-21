package leetcode

import (
	"fmt"
	"sort"
)

type NumFrequency struct {
	F int
	N int
}

func RunTopK() {
	testcase := []int{11, 11, 11, 22, 22, 22, 22, 44, 44, 44, 44, 44, 44, 33}
	fmt.Println(topKFrequent(testcase, 2))
}

func topKFrequent(nums []int, k int) []int {
	freqs := make(map[int]NumFrequency)

	for _, num := range nums {
		if _, exists := freqs[num]; !exists {
			freqs[num] = NumFrequency{
				F: 1,
				N: num,
			}
		} else {
			freq := freqs[num].F + 1
			newV := NumFrequency{
				F: freq,
				N: num,
			}
			freqs[num] = newV
		}
	}

	freqsSlice := []NumFrequency{}
	for _, freq := range freqs {
		freqsSlice = append(freqsSlice, freq)
	}

	sort.Slice(freqsSlice, func(i, j int) bool {
		return freqsSlice[i].F > freqsSlice[j].F
	})

	res := []int{}
	for _, sl := range freqsSlice[:k] {
		res = append(res, sl.N)
	}
	return res
}
