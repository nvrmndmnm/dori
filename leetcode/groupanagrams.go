package leetcode

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func Run() {
	body, err := os.ReadFile("leetcode/inputs/groupanagrams")
	if err != nil {
		fmt.Printf("unable to read file: %v", err)
	}
	testcase := strings.Split(string(body), ",")
	fmt.Println(groupAnagrams(testcase))
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	group := make(map[string][]string)

	for _, str := range strs {
		key := sortFunc(str)
		if sortFunc(str) == key {
			group[key] = append(group[key], str)
		}
	}

	for _, v := range group {
		res = append(res, v)
	}
	return res
}

func sortFunc(str string) string {
	runeSlice := []rune(str)
	sort.Slice(runeSlice, func(i int, j int) bool {
		return runeSlice[i] < runeSlice[j]
	})
	return string(runeSlice)
}
