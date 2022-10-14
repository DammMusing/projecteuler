/**
Problem 32: Pandigital products

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
*/

package main

func SumOfFoundPandigitalProducts(limit int) int {
	var sum int
	pandigitals := make(map[int]bool)

	for i := 1; i < limit; i++ {
		for j := i + 1; j < limit; j++ {
			product := i * j
			if is_pandigital(i, j, product) {
				if !pandigitals[product] {
					pandigitals[product] = true
					sum += product
				}
			}
		}
	}

	return sum
}

func is_pandigital(i, j, p int) bool {
	digits := make(map[int]int)

	for i > 0 {
		digit := i % 10
		if digit == 0 || digits[digit] > 0 {
			// Only interested in 1-9 pandigital values
			// and can short-circuit on any duplicates.
			return false
		}
		digits[digit] = 1
		i /= 10
	}
	for j > 0 {
		digit := j % 10
		if digit == 0 || digits[digit] > 0 {
			return false
		}
		digits[digit] = 1
		j /= 10
	}
	for p > 0 {
		digit := p % 10
		if digit == 0 || digits[digit] > 0 {
			return false
		}
		digits[digit] = 1
		p /= 10
	}

	// No duplicates, but make sure each digit is represented.
	return len(digits) == 9
}
