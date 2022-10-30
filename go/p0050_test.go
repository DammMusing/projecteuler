/**
Problem 50: Consecutive prime sum

The prime 41, can be written as the sum of six consecutive primes:

41 = 2 + 3 + 5 + 7 + 11 + 13
This is the longest sum of consecutive primes that adds to a prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most consecutive primes?
*/

package main

import "testing"

func TestSumOfMostConsecutivePrimesAndPrime(t *testing.T) {
	type args struct {
		limit uint64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "small", args: args{limit: 100}, want: 41},
		{name: "thousand", args: args{limit: 1000}, want: 953},
		{name: "answer", args: args{limit: 1000000}, want: 997651},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfMostConsecutivePrimesAndPrime(tt.args.limit); got != tt.want {
				t.Errorf("SumOfMostConsecutivePrimesAndPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
