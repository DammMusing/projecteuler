/**
Problem 30: Digit fifth powers

Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

1634 = 14 + 64 + 34 + 44
8208 = 84 + 24 + 04 + 84
9474 = 94 + 44 + 74 + 44
As 1 = 14 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/

package main

import "strconv"

// I thought I might have to build a map of the totally-ordered digits
// to cut some corners in the search space, but a brute force approach
// finishes in less than 2.3 seconds.  Go figure!
func SumOfDigitPowersWithExponent(exp int) int {
	total := 0

	powers := make([]int, 10)
	for i := range powers {
		powers[i] = 1
		for e := 0; e < exp; e++ {
			powers[i] *= i
		}
	}

	for n := 10; n < 1000000; n++ {
		sum := 0
		digits := strconv.Itoa(n)
		for _, digit := range digits {
			sum += powers[int(digit-'0')]
		}
		if sum == n {
			total += sum
		}
	}

	return total
}
