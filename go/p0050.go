/**
Problem 50: Consecutive prime sum

The prime 41, can be written as the sum of six consecutive primes:

41 = 2 + 3 + 5 + 7 + 11 + 13
This is the longest sum of consecutive primes that adds to a prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most consecutive primes?
*/

package main

func SumOfMostConsecutivePrimesAndPrime(limit uint64) int {
	prime_list, prime_map := make([]int, 0), make(map[int]bool)
	prime_chan := make(chan uint64)
	go GatherPrimesLessThan(limit, prime_chan)

	for prime := <-prime_chan; prime != 0; prime = <-prime_chan {
		prime_list = append(prime_list, int(prime))
		prime_map[int(prime)] = true
	}

	max_consecutive := 0
	biggest_prime := 0
	for i := 0; i < int(limit); i++ {
		sum := 0
		for j := 0; i+j < len(prime_list); j++ {
			sum += prime_list[i+j]
			if prime_map[sum] && sum < int(limit) && j > max_consecutive {
				max_consecutive = j
				biggest_prime = sum
			}
		}
	}

	return biggest_prime
}
