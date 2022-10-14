/**
Problem 32: Pandigital products

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
*/

package main

import (
	"testing"
)

func TestSumOfFoundPandigitalProducts(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{limit: 140}, want: 5796},
		{name: "answer", args: args{limit: 10000}, want: 45228},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfFoundPandigitalProducts(tt.args.limit); got != tt.want {
				t.Errorf("SumOfFoundPandigitalProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_pandigital(t *testing.T) {
	type args struct {
		i int
		j int
		p int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "sample", args: args{i: 39, j: 186, p: 7254}, want: true},
		{name: "badzample", args: args{i: 43, j: 789, p: 41265}, want: false},
		{name: "weirdzample", args: args{i: 43, j: 789, p: 1256}, want: true},
		{name: "incompletzample", args: args{i: 2, j: 3, p: 6}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := is_pandigital(tt.args.i, tt.args.j, tt.args.p); got != tt.want {
				t.Errorf("is_pandigital() = %v, want %v", got, tt.want)
			}
		})
	}
}
