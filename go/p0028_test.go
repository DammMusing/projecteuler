/**
Problem 28: Number sprial diagonals

Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:

21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
*/

package main

import "testing"

func TestSumDiagonalsOnSpiralOfSize(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "example", args: args{size: 5}, want: 101},
		{name: "answer", args: args{size: 1001}, want: 669171001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumDiagonalsOnSpiralOfSize(tt.args.size); got != tt.want {
				t.Errorf("SumDiagonalsOnSpiralOfSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
