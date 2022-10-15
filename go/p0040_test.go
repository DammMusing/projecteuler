/**
Problem 40: Champernowne's constant

An irrational decimal fraction is created by concatenating the positive integers:

    0.123456789101112131415161718192021...
		             ^
                 12th digit

It can be seen that the 12th digit of the fractional part is 1.

If d_n represents the nth digit of the fractional part, find the value of the following expression.

d_1 × d_10 × d_100 × d_1000 × d_10000 × d_100000 × d_1000000
*/

package main

import "testing"

func TestChampernowneConstant(t *testing.T) {
	type args struct {
		lg_digits int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{lg_digits: 2}, want: 1},
		{name: "simple2", args: args{lg_digits: 3}, want: 5},
		{name: "answer", args: args{lg_digits: 7}, want: 210},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChampernowneConstant(tt.args.lg_digits); got != tt.want {
				t.Errorf("ChampernowneConstant() = %v, want %v", got, tt.want)
			}
		})
	}
}
