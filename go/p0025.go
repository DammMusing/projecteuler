/**
Problem 25: 1000 digit Fibonacci number

The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
*/

package main

import "math/big"

// Returns the index of the first element of the sequence having so many digits.
func FindFirstFibonacciOfNDigits(ndigits int) int {
	fib1, fib2 := big.NewInt(1), big.NewInt(1)
	index := 2
	for ; len(fib2.String()) < ndigits; index++ {
		sum := new(big.Int).Set(fib2)
		sum.Add(sum, fib1)
		fib1 = fib2
		fib2 = sum
	}
	return index
}
