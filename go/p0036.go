/**
Problem 36: Double-base palindromes

The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)
*/

package main

import "math/big"

func SumOfDoubleBasePalindromicUpTo(limit int) int {
	sum := 0

	for i := 1; i < limit; i++ {
		value := big.NewInt(int64(i))
		if is_palindrome_string(value.Text(10)) && is_palindrome_string(value.Text(2)) {
			sum += int(value.Int64())
		}
	}
	return sum
}

func is_palindrome_string(value string) bool {
	for i := 0; i < len(value)/2+1; i++ {
		if value[i] != value[len(value)-i-1] {
			return false
		}
	}
	return true
}
