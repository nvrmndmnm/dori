package leetcode

import "fmt"

func RunPalindrome() {
	// testcase := ""
	// testcase := "A man, a plan, a canal: Panama"
	testcase := "0P"
	fmt.Println(isPalindrome(testcase))
}

func isPalindrome(s string) bool {
	runeArr := make([]rune, 0)
	for _, c := range s {
		fmt.Println(c)
		if c >= 65 && c <= 90 {
			runeArr = append(runeArr, c+32)
		}
		if c >= 97 && c <= 122 || c >= 48 && c <= 57 {
			runeArr = append(runeArr, c)
		}
	}

	if len(runeArr) == 0 {
		return true
	}
	counter := 0
	for i := len(runeArr) - 1; i > 0; i-- {
		if runeArr[i] == runeArr[len(runeArr)-1-i] {
			counter++
		}
	}

	return counter == len(runeArr)-1
}
