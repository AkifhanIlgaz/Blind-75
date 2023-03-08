package main

import (
	"unicode"
)

func main() {

}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	arr := []rune(s)

	for left < right {
		leftChar := unicode.ToLower(arr[left])
		rightChar := unicode.ToLower(arr[right])

		if !isAlphaNumeric(leftChar) {
			left++
			continue
		}

		if !isAlphaNumeric(rightChar) {
			right--
			continue
		}

		if leftChar != rightChar {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNumeric(s rune) bool {
	return unicode.IsDigit(s) || unicode.IsLetter(s)
}
