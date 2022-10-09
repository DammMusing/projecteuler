/**
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

package main

func NthPrime(n int) uint64 {
	var nth_prime uint64 = 1
	primes := make(chan uint64)
	go GatherPrimesLessThan(1000000, primes)

	i := 0
	for prime := <-primes; prime != 0; prime = <-primes {
		nth_prime = prime
		i++
		if i == n {
			break
		}
	}

	return nth_prime
}
