/**
Problem 26: Reciprocal cycles

A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

1/2	= 	0.5
1/3	= 	0.(3)
1/4	= 	0.25
1/5	= 	0.2
1/6	= 	0.1(6)
1/7	= 	0.(142857)
1/8	= 	0.125
1/9	= 	0.(1)
1/10	= 	0.1
Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
*/

package main

// Let's relearn long division!  plus some pattern detection

func LongestRecurringReciprocalCycle(limit int) int {
	longest, answer := 0, 0

	for d := 1; d < limit; d++ {
		digits := long_div_pattern(d)
		cycle_length := len(digits)
		if cycle_length > longest {
			longest = cycle_length
			answer = d
		}
	}

	return answer
}

// Returns the recurring cycle (may omit leading digits)
func long_div_pattern(divisor int) []int {
	digits := make([]int, 0, 32)
	remainders := make([]int, 1, 32)
	remainder := 10
	remainders[0] = remainder

	for remainder != 0 {
		if remainder < divisor {
			remainder *= 10
			digits = append(digits, 0)
			remainders = append(remainders, remainder)
			continue
		}

		quotient := int(remainder / divisor)
		produced := quotient * divisor
		digits = append(digits, digits_of(quotient)...)
		pattern_length := pattern_find(remainders)
		if pattern_length > 0 {
			return digits[len(digits)-pattern_length:]
		}

		remainder = (remainder - produced) * 10
		remainders = append(remainders, remainder)
	}

	// We found a perfect multiple, return 'no pattern'
	return nil
}

func digits_of(value int) []int {
	digits := make([]int, 0, 8)
	for value >= 10 {
		digits = append(digits, value%10)
		value /= 10
	}
	digits = append(digits, value)
	// Reverse this list to get expected decreasing ordering.
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}

func pattern_find(values []int) int {
	// Try gradually increasing sizes of pattern windows, testing for validity.
	for pattern_length := 1; pattern_length <= len(values)/3; pattern_length++ {
		found := true
		for i := 0; i < pattern_length; i++ {
			e := len(values)
			a, b, c := e-pattern_length, e-pattern_length*2, e-pattern_length*3
			if values[a] != values[b] || values[a] != values[c] {
				found = false
				break
			}
		}
		if found {
			return pattern_length
		}
	}

	return 0
}
