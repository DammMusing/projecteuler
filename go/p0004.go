/**
Problem 4: Largest Palindrome Product

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import "strconv"

func is_palindrome(number int) bool {
	digits := strconv.Itoa(number)
	for i := 0; i < len(digits)/2+1; i++ {
		if digits[i] != digits[len(digits)-i-1] {
			return false
		}
	}
	return true
}

func LargestPalindromeProduct(multiplicand_limit int) int {
	var largest int

	for i := 2; i < multiplicand_limit; i++ {
		for j := i; j < multiplicand_limit; j++ {
			multiple := i * j
			if is_palindrome(multiple) && multiple > largest {
				largest = multiple
			}
		}
	}

	return largest
}
