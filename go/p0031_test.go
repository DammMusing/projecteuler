/**
Problem 31: Coin sums

In the United Kingdom the currency is made up of pound (£) and pence (p). There are eight coins in general circulation:

1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), and £2 (200p).
It is possible to make £2 in the following way:

1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?
*/

package main

import "testing"

func TestCountWaysToMakeTwoPoundsSterling(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "answer", want: 73682},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountWaysToMakeTwoPoundsSterling(); got != tt.want {
				t.Errorf("CountWaysToMakeTwoPoundsSterling() = %v, want %v", got, tt.want)
			}
		})
	}
}
