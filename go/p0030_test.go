/**
Problem 30: Digit fifth powers

Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

1634 = 14 + 64 + 34 + 44
8208 = 84 + 24 + 04 + 84
9474 = 94 + 44 + 74 + 44
As 1 = 14 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/

package main

import "testing"

func TestSumOfDigitPowersWithExponent(t *testing.T) {
	type args struct {
		exp int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{exp: 4}, want: 19316},
		{name: "answer", args: args{exp: 5}, want: 443839},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfDigitPowersWithExponent(tt.args.exp); got != tt.want {
				t.Errorf("SumOfDigitPowersWithExponent() = %v, want %v", got, tt.want)
			}
		})
	}
}
