/**
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

package main

import "testing"

func TestNthPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "simple", args: args{n: 6}, want: 13},
		{name: "simple", args: args{n: 10001}, want: 104743},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NthPrime(tt.args.n); got != tt.want {
				t.Errorf("NthPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
