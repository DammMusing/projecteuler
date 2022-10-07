/**
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import (
	"github.com/DammMusing/projecteuler/utils"
)

func SumOfPrimesBelow(n uint64) uint64 {
	summation := uint64(0)
	primes := make(chan uint64)
	go utils.GatherPrimesLessThan(n, primes)

	for prime := <-primes; prime != 0; prime = <-primes {
		if prime < n {
			summation += prime
		}
	}

	return summation
}
