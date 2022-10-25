/**
Problem 38: Pandigital multiples

Take the number 192 and multiply it by each of 1, 2, and 3:

192 × 1 = 192
192 × 2 = 384
192 × 3 = 576

By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?
*/

package main

import "fmt"

func LargestPandigitalFromConcatenatedSerialProduct(limit int) int {
	max_pandigital := 0

	for i := 1; i < limit; i++ {
		i_pan := product_9digit(i)
		if !is_pandigital_int(i_pan) {
			continue
		}
		fmt.Printf("found %d: %d\n", i, i_pan)
		if i_pan > max_pandigital {
			max_pandigital = i_pan
		}
	}

	return max_pandigital
}

func product_9digit(gen int) int {
	palindrome := 0
	for i := 1; palindrome < 100000000; i++ {
		x := gen * i
		shift := round_pow10(x)
		palindrome = (palindrome * shift) + x
	}
	return palindrome
}

// Returns the smallest power of ten greater than value.
func round_pow10(value int) int {
	pow := 1
	for pow <= value {
		pow *= 10
	}
	return pow
}

func is_pandigital_int(value int) bool {
	if value < 100000000 || value >= 1000000000 {
		// Must be a nine-digit number
		return false
	}
	digits := make(map[int]bool)
	for value > 0 {
		digit := value % 10
		if digit == 0 || digits[digit] {
			// disallow zeroes and any second occurrence of a digit
			return false
		}
		digits[digit] = true
		value /= 10
	}
	return len(digits) == 9
}
