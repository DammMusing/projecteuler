/**
The sequence of triangle numbers is generated by adding the natural numbers. So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28. The first ten terms would be:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

Let us list the factors of the first seven triangle numbers:

 1: 1
 3: 1,3
 6: 1,2,3,6
10: 1,2,5,10
15: 1,3,5,15
21: 1,3,7,21
28: 1,2,4,7,14,28
We can see that 28 is the first triangle number to have over five divisors.

What is the value of the first triangle number to have over five hundred divisors?
*/

package main

import (
	"github.com/DammMusing/projecteuler/utils"
)

// Brute-force this takes some time but still O(minutes).  A more efficient
// solution would generate the factors from the bottom up instead of finding
// factors (linear time) while iterating in the domain (making it quadratic).
// The generated approach would also be quadratic but with much lower constant.
func FirstTriangleNumberOfManyFactors(min_num_divisors int) uint64 {
	triangle, sequence := uint64(1), uint64(2)
	for num_factors := 1; num_factors < min_num_divisors; sequence++ {
		triangle += sequence
		num_factors = utils.NumFactorsOf(triangle)
	}

	return triangle
}
