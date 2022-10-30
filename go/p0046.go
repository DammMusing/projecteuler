/**
Problem 46: Goldbach's other conjecture

It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.

9 = 7 + 2×12
15 = 7 + 2×22
21 = 3 + 2×32
25 = 7 + 2×32
27 = 19 + 2×22
33 = 31 + 2×12

It turns out that the conjecture was false.

What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?
*/

package main

func FirstCompositeNotAPrimePlusDoubleSquare(limit int) int {
	primes := make_prime_map(limit)
	evidence := make(map[int]bool)
	for prime := range primes {
		for i, v := 1, prime+2; v < limit; i, v = i+1, prime+2*(i+1)*(i+1) {
			evidence[v] = true
		}
	}

	odd_composites := make(chan int)
	go GatherOddComposites(limit, primes, odd_composites)

	for number := <-odd_composites; number != 0; number = <-odd_composites {
		if !evidence[number] {
			return number
		}
	}

	return 0
}

func make_prime_map(limit int) map[int]bool {
	primes := make(map[int]bool)
	prime_chan := make(chan uint64)
	go GatherPrimesLessThan(uint64(limit), prime_chan)

	for prime := <-prime_chan; prime != 0; prime = <-prime_chan {
		primes[int(prime)] = true
	}
	return primes
}

func GatherOddComposites(limit int, primes map[int]bool, out chan int) {
	for i := 3; i < limit; i += 2 {
		if primes[i] {
			continue
		}
		out <- i
	}
	out <- 0
}
