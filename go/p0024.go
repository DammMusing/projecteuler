/**
Problem 24: Lexicographic permutations

A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

package main

// We can take a shortcut here by knowing that the answer begins with '2'
// because there are 9! = 362880 values with the remaining 9 digits when
// choosing a leading '0', another 362880 when with a leading '1'.  We can
// follow this reasoning to determine the remaining digits' placement.

func MillionthUniqueDigitPermutation(place int) int {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	output := []int{}

	// select successively larger digits, staying below the (remaining) placement.
	for len(digits) > 0 {
		increment := int_factorial(len(digits) - 1)
		for i, digit := range digits {
			if place <= increment {
				output = append(output, digit)
				digits = append(digits[:i], digits[i+1:]...)
				break
			}
			place -= increment
		}
	}

	result := 0
	for _, v := range output {
		result = result*10 + v
	}
	return result
}

// It's not the most efficient factorial, nor is it the most elegant, but we
// know we'll be under 10,000,000 and only call this a few times per digit.
func int_factorial(n int) int {
	if n <= 1 {
		return 1
	}
	agg := n
	for mul := n - 1; mul > 1; mul-- {
		agg *= mul
	}
	return agg
}
