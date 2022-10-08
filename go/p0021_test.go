/**
Problem 21: Amicable numbers

Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

package main

import "testing"

func TestSumAmicableNumbersUpTo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{n: 1000}, want: 504}, // = { 220, 284 }
		{name: "answer", args: args{n: 10000}, want: 31626},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAmicableNumbersUpTo(tt.args.n); got != tt.want {
				t.Errorf("SumAmicableNumbersUpTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
