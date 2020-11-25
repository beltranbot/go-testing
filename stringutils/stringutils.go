package stringutils

import "strings"

// ana
// ivi
// anita lava la tina

// IsPalindrome func
func IsPalindrome(text string) bool {
	return isPalindrome(strings.Replace(text, " ", "", -1))
}

func isPalindrome(text string) bool {
	if len(text) <= 1 { // "a" ""
		return true
	}

	if len(text) == 2 && text[0] != text[1] { // "aa"
		return false
	}

	if text[0] != text[len(text)-1] { // "ana"
		return false
	}

	return isPalindrome(text[1 : len(text)-1])
}
