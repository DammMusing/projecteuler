/**
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
(note: a + b + c = 12 here)

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package main

const LIMIT int = 1000

func PythagoreanTripletWithSum(n int) int {
	SQUARES := make(map[int]int)
	for i := 0; i < LIMIT; i++ {
		SQUARES[i*i] = i
	}

	for a := 3; a < LIMIT; a++ {
		for b := a + 1; b < LIMIT; b++ {
			c2 := (a * a) + (b * b)
			c := SQUARES[c2]
			if c == 0 {
				continue
			}
			if a+b+c == n {
				return a * b * c
			}
		}
	}
	// sentinel; halt only if a condition is met above or we exhaust the domain
	return 0
}
