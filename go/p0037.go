/**
Problem 37: Truncatable primes

The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
*/

package main

import "math/big"

func SumOfAllTruncatablePrimes() uint64 {
	count := 11
	sum := uint64(0)

	primes := make(map[uint64]bool)
	prime_chan := make(chan uint64)
	go GatherPrimesLessThan(1000000, prime_chan)
	for prime := <-prime_chan; prime != uint64(0); prime = <-prime_chan {
		primes[prime] = true
		if prime < 10 {
			continue
		}
		if is_truncatable(prime, primes) {
			count -= 1
			sum += prime
		}
		if count == 0 {
			break
		}
	}

	return sum
}

func is_truncatable(prime uint64, primes map[uint64]bool) bool {
	digits := make([]int, 0)
	for _, digit := range big.NewInt(int64(prime)).Text(10) {
		digits = append(digits, int(digit-'0'))
	}

	// Left truncations
	for i := 0; i < len(digits); i++ {
		truncated := from_digits(digits[i:])
		if truncated == 0 || truncated == 1 {
			return false
		}
		if !primes[uint64(truncated)] {
			return false
		}
	}
	// Right truncations
	for j := len(digits); j > 0; j-- {
		truncated := from_digits(digits[:j])
		if truncated == 0 || truncated == 1 {
			return false
		}
		if !primes[uint64(truncated)] {
			return false
		}
	}

	return true
}
