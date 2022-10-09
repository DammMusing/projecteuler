/**
Problem 24: Lexicographic permutations

A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

package main

import "testing"

func TestMillionthUniqueDigitPermutation(t *testing.T) {
	type args struct {
		place int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "first", args: args{place: 1}, want: 123456789},
		{name: "simple", args: args{place: 3}, want: 123456879},
		{name: "answer", args: args{place: 1000000}, want: 2783915460},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MillionthUniqueDigitPermutation(tt.args.place); got != tt.want {
				t.Errorf("MillionthUniqueDigitPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
