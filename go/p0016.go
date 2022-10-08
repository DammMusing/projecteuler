/**
Problem 16: Power digit sum

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

package main

import "math/big"

func SumOfDigitsOfPowerOfTwo(power int64) int {
	sum := 0
	num, pow := big.NewInt(2), big.NewInt(power)
	num.Exp(num, pow, nil)

	digits := num.Text(10)
	for _, x := range digits {
		sum += int(x - '0')
	}

	return sum
}
