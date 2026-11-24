package main

import "fmt"

// IsPalindrome reports whether s reads the same both ways once everything
// but letters and digits is ignored and case is folded. Two pointers that
// skip non-alphanumerics in place — no cleaned copy of the string.
func IsPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !alnum(s[l]) {
			l++
		}
		for l < r && !alnum(s[r]) {
			r--
		}
		if lower(s[l]) != lower(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func alnum(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9'
}

func lower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 'a' - 'A'
	}
	return c
}

func main() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("race a car"))
}
