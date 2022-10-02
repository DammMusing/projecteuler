/**
Problem 1: Multiples of 3 or 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

func DoSum(upper_bound int) int {
	var sum int

	for i := 5; i < upper_bound; i += 5 {
		sum += i
	}
	for i := 3; i <= upper_bound; i += 3 {
		// Add only new values that were not already added in above loop.
		if i%5 == 0 {
			continue
		}
		sum += i
	}

	return sum
}
