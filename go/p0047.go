/**
Problem 47: Distinct primes factors

The first two consecutive numbers to have two distinct prime factors are:

14 = 2 × 7
15 = 3 × 5

The first three consecutive numbers to have three distinct prime factors are:

644 = 2² × 7 × 23
645 = 3 × 5 × 43
646 = 2 × 17 × 19.

Find the first four consecutive integers to have four distinct prime factors each. What is the first of these numbers?
*/

package main

func ConsecutiveIntegersWithDistinctPrimes(count int) []int {
	candidates := make([]int, count)
	primes := make_prime_list(1000000)
	for i := 0; i < count; i++ {
		candidates[i] = i + 1
	}

	for !distinct_prime_factors(candidates, primes) {
		// slide the candidates window forward
		for i, v := range candidates {
			candidates[i] = v + 1
		}
	}
	return candidates
}

func make_prime_list(limit int) []int {
	primes := make([]int, 0)
	prime_chan := make(chan uint64)
	go GatherPrimesLessThan(uint64(limit), prime_chan)

	for prime := <-prime_chan; prime != 0; prime = <-prime_chan {
		primes = append(primes, int(prime))
	}
	return primes
}

func distinct_prime_factors(values []int, primes []int) bool {
	for _, value := range values {
		count_factors := 0
		for _, factor := range primes {
			if factor > value {
				break
			}
			if value%factor == 0 {
				count_factors += 1
				for value%factor == 0 {
					value /= factor
				}
			}
			if value <= 1 {
				break
			}
		}

		if count_factors != len(values) {
			return false
		}
	}
	return true
}
