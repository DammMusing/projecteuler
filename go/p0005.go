/**
Problem 5: Smallest multiple

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

func divisible_by_range(number, limit int) bool {
	for i := 2; i <= limit; i++ {
		if number%i != 0 {
			return false
		}
	}
	return true
}

func SmallestMultiple(limit int) int {
	candidate := limit
	for !divisible_by_range(candidate, limit) {
		candidate += limit
	}
	return candidate
}
