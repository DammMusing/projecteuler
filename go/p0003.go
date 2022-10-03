/**
Problem 3: Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import (
	"math"
)

// Not the most efficient primality check but good enough for this problem.
func is_prime(number int) bool {
	limit := int(math.Sqrt(float64(number))) + 1
	for i := 2; i < limit; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func LargestPrimeFactor(number int) int {
	limit := int(math.Sqrt(float64(number))) + 1
	var max_factor int
	for i := 2; i < limit; i++ {
		if number%i == 0 {
			if is_prime(i) {
				max_factor = i
			}
		}
	}
	return max_factor
}
