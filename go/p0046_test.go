/**
Problem 46: Goldbach's other conjecture

It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.

9 = 7 + 2×12
15 = 7 + 2×22
21 = 3 + 2×32
25 = 7 + 2×32
27 = 19 + 2×22
33 = 31 + 2×12

It turns out that the conjecture was false.

What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?
*/

package main

import "testing"

func TestFirstCompositeNotAPrimePlusDoubleSquare(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{limit: 1000}, want: 0},
		{name: "answer", args: args{limit: 10000}, want: 5777},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstCompositeNotAPrimePlusDoubleSquare(tt.args.limit); got != tt.want {
				t.Errorf("FirstCompositeNotAPrimePlusDoubleSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
