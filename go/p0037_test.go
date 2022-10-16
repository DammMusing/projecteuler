/**
Problem 37: Truncatable primes

The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
*/

package main

import (
	"testing"
)

func Test_is_truncatable(t *testing.T) {
	type args struct {
		prime uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "yes", args: args{prime: 3797}, want: true},
		{name: "no", args: args{prime: 973}, want: false},
	}
	primes := make(map[uint64]bool)
	primes[3797] = true
	primes[379] = true
	primes[37] = true
	primes[3] = true
	primes[797] = true
	primes[973] = true
	primes[97] = true
	primes[9] = true
	primes[7] = true
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := is_truncatable(tt.args.prime, primes); got != tt.want {
				t.Errorf("is_truncatable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumOfAllTruncatablePrimes(t *testing.T) {
	tests := []struct {
		name string
		want uint64
	}{
		{name: "answer", want: 748317},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfAllTruncatablePrimes(); got != tt.want {
				t.Errorf("SumOfAllTruncatablePrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
