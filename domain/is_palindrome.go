package domain

// IsPalindrome returns true if s is a palindrome.
func IsPalindrome(s string) bool {
	return s == Reverse(s)
}
