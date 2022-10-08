/**
Problem 20: Factorial digit sum

n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

package main

import "math/big"

func SumOfDigitsFromFactorial(n int64) int {
	sum_digits := 0
	digits := factorial(n).Text(10)
	for _, v := range digits {
		sum_digits += int(v - '0')
	}
	return sum_digits
}

func factorial(n int64) *big.Int {
	if n <= 1 {
		return big.NewInt(1)
	}
	agg := new(big.Int).SetInt64(n)
	for mul := n - 1; mul > 1; mul-- {
		agg.Mul(agg, big.NewInt(mul))
	}
	return agg
}
