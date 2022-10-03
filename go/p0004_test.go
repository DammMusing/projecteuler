/**
Problem 4: Largest Palindrome Product

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import "testing"

func TestLargestPalindromeProduct(t *testing.T) {
	type args struct {
		multiplicand_limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{multiplicand_limit: 100}, want: 9009},
		{name: "answer", args: args{multiplicand_limit: 1000}, want: 906609},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestPalindromeProduct(tt.args.multiplicand_limit); got != tt.want {
				t.Errorf("LargestPalindromeProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
