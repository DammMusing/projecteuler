/**
Problem 27: Quadratic primes

Euler discovered the remarkable quadratic formula:

n^2 + n + 41

It turns out that the formula will produce 40 primes for the consecutive integer values . However, when n=40, the result is divisible by 41, and certainly when n=41 it is clearly divisible by 41.

The incredible formula n^2 - 79n + 1601 was discovered, which produces 80 primes for the consecutive values 0 <= n <= 79. The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

   n^2 + an + b, where |a| < 1000 and |b| <= 1000

   where |n| is modulus/absolute value of n
	 e.g. |11| = 11 and |-4| = 4

Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for consecutive values of n, starting with n=0.
*/

package main

func ProductOfCoefficientsForQuadraticPrimeFormula(limit int) int {
	max_count := 0
	prod_ab := 0

	// Find enough primes to cover codomain.
	primes := make(map[int]bool)
	primechan := make(chan uint64)
	go GatherPrimesLessThan(uint64(2*limit*limit), primechan)
	for prime := <-primechan; prime != 0; prime = <-primechan {
		primes[int(prime)] = true
	}

	for a := -limit + 1; a < limit; a++ {
		for b := -limit; b <= limit; b++ {
			num_primes := count_primes_from_quadratic(a, b, primes)
			if num_primes > max_count {
				max_count = num_primes
				prod_ab = a * b
			}
		}
	}

	return prod_ab
}

func count_primes_from_quadratic(a int, b int, primes map[int]bool) int {
	quad := func(n, a, b int) int {
		return (n * n) + (a * n) + b
	}

	count := 0
	for ; primes[quad(count, a, b)]; count++ {
	}

	return count
}
