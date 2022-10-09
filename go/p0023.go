/**
Problem 23: Non-abundant sums

A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/

package main

func SumOfNonAbundantSums(limit int) uint64 {
	// Build the sum of divisors mapping from the bottom up.
	sum_divisors := make([]int, limit)
	sum_divisors[0] = 0
	sum_divisors[1] = 1
	for d := 2; d < limit; d++ {
		sum_divisors[d] += 1
		for id := d * 2; id < limit; id += d {
			sum_divisors[id] += d
		}
	}

	abundants := make([]int, 0, 1000)
	for i, v := range sum_divisors {
		if i < v {
			abundants = append(abundants, i)
		}
	}

	// Now it's somewhat convenient that we didn't bake any primality logic into the sieve
	// ... but perhaps this and that deserve to be different types (Eratosthenes vs Bitmap).
	sieve := NewSieve(uint64(limit))
	for i := 0; i < len(abundants); i++ {
		for j := i; j < len(abundants); j++ {
			sum := abundants[i] + abundants[j]
			sieve.Set(uint64(sum), true)
		}
	}

	total := uint64(0)
	for i := uint64(0); i < uint64(limit); i++ {
		if !sieve.Get(i) {
			total += i
		}
	}
	return total
}
