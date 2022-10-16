/**
Problem 36: Double-base palindromes

The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)
*/

package main

import "testing"

func TestSumOfDoubleBasePalindromicUpTo(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{limit: 700}, want: 1055},
		{name: "simple", args: args{limit: 1000000}, want: 872187},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfDoubleBasePalindromicUpTo(tt.args.limit); got != tt.want {
				t.Errorf("SumOfDoubleBasePalindromicUpTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
