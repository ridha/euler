// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 x 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
	"strconv"
)

func is_palindrome(word string) bool {
	len_word := len(word)
	for i := 0; i < len_word/2; i++ {
		if word[i] != word[len_word-i-1] {
			return false
		}
	}
	return true
}

func main() {
	max := 0
	for a := 999; a > 900; a-- {
		for b := 999; b > 900; b-- {
			if is_palindrome(strconv.Itoa(a * b)) {
				if a*b > max {
					max = a * b
				}
			}
		}
	}
	fmt.Println(max)
}
