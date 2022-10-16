/**
Problem 34: Digit factorials

145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: As 1! = 1 and 2! = 2 are not sums they are not included.
*/

package main

import "testing"

func TestSumOfNumbersEqualToFactorialOfTheirDigits(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{limit: 150}, want: 145},
		{name: "answer", args: args{limit: 50000}, want: 40730},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfNumbersEqualToFactorialOfTheirDigits(tt.args.limit); got != tt.want {
				t.Errorf("SumOfNumbersEqualToFactorialOfTheirDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
