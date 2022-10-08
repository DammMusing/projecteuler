/**
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

package main

import "testing"

func TestLongestCollatzSequenceBelow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{n: 7}, want: 6},
		{name: "simple+", args: args{n: 9}, want: 7},
		{name: "simple+2", args: args{n: 10}, want: 9},
		{name: "answer", args: args{n: 1000000}, want: 837799},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCollatzSequenceBelow(tt.args.n); got != tt.want {
				t.Errorf("LongestCollatzSequenceBelow() = %v, want %v", got, tt.want)
			}
		})
	}
}
