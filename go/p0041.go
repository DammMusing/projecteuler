/**
Problem 41: Pandigital prime

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and is also prime.

What is the largest n-digit pandigital prime that exists?
*/

package main

func LargestPandigitalPrime() int {
	max_pp := 0

	// We could look up through all 9 digit numbers but it's not a full nine-
	// pandigital.. so we can finish this faster instead.
	prime_chan := make(chan uint64)
	go GatherPrimesLessThan(uint64(100000000), prime_chan)
	for p := <-prime_chan; p != 0; p = <-prime_chan {
		pint := int(p)
		if !is_pandigital_int(pint) {
			continue
		}
		if pint > max_pp {
			max_pp = pint
		}
	}

	return max_pp
}
