package domain

import "strings"

// IsPalindrome returns true if s is a palindrome.
func IsPalindrome(s string) bool {
	return s == Reverse(s)
}
