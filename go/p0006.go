/**
The sum of the squares of the first ten natural numbers is,

   1^2 + 2^2 + 3^2 ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,

   (1 + 2 + 3 ... + 10)^2 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is .

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

package main

import "math/big"

// Will brute force work for this just because of BigInt?.. or will we have to play with factors
func SumSquareDifference(limit int64) big.Int {
	var sum_squares, square_sum big.Int
	for i := int64(1); i <= limit; i++ {
		bigi := *big.NewInt(i)
		square_sum.Add(&square_sum, &bigi)
		sum_squares.Add(&sum_squares, bigi.Mul(&bigi, &bigi))
	}
	square_sum.Mul(&square_sum, &square_sum)

	diff := square_sum.Sub(&square_sum, &sum_squares)
	return *diff
}
