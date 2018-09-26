// Package main provides utilities for word games....
package main

import (
	"fmt"
)

// IsPalindrome reports whether s reads the same forward and backward.
// My first attempt!

func main() {
	fmt.Printf("%t\n", IsPalindrome("Tony"))
}

// IsPalindrome function
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
