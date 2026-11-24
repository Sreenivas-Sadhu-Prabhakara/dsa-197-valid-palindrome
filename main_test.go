package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true}, // nothing alphanumeric: empty reads the same both ways
		{"", true},
		{"0P", false}, // digit vs letter, the classic trap
		{"ab_a", true},
		{"Madam", true},
		{"12321", true},
	}
	for _, c := range cases {
		if got := IsPalindrome(c.in); got != c.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}
