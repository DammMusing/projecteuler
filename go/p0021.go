/**
Problem 21: Amicable numbers

Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

package main

func SumAmicableNumbersUpTo(n int) int {
	sum_all := 0

	// Build the sum of divisors mapping from the bottom up.
	sum_divisors := make([]int, n)
	sum_divisors[1] = 1
	for d := 2; d < n; d++ {
		sum_divisors[d] += 1
		for id := d * 2; id < n; id += d {
			sum_divisors[id] += d
		}
	}

	for i, divsum := range sum_divisors {
		if divsum < n && sum_divisors[divsum] == i && i != divsum {
			sum_all += divsum
		}
	}
	return sum_all
}
