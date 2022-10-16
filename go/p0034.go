/**
Problem 34: Digit factorials

145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: As 1! = 1 and 2! = 2 are not sums they are not included.
*/

package main

func SumOfNumbersEqualToFactorialOfTheirDigits(limit int) int {
	sum_total := 0
	factorials := make_factorials(10)

	for x := 3; x < limit; x++ {
		digits := digits_of(x)
		sum := 0
		for _, digit := range digits {
			sum += factorials[digit]
		}
		if sum == x {
			sum_total += sum
		}
	}

	return sum_total
}

func make_factorials(limit int) []int {
	facts := make([]int, 10)
	facts[0] = 1
	for i := 1; i < limit; i++ {
		facts[i] = facts[i-1] * i
	}
	return facts
}
