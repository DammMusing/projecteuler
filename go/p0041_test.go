/**
Problem 41: Pandigital prime

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and is also prime.

What is the largest n-digit pandigital prime that exists?
*/

package main

import "testing"

func TestLargestPandigitalPrime(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "answer", want: 7652413},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestPandigitalPrime(); got != tt.want {
				t.Errorf("LargestPandigitalPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
