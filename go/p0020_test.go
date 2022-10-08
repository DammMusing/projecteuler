/**
Problem 20: Factorial digit sum

n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

package main

import "testing"

func TestSumOfDigitsFromFactorial(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{n: 10}, want: 27},
		{name: "simple", args: args{n: 100}, want: 648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfDigitsFromFactorial(tt.args.n); got != tt.want {
				t.Errorf("SumOfDigitsFromFactorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
