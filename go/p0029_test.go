/**
Problem 29: Distinct powers

Consider all integer combinations of a^b for 2 ≤ a ≤ 5 and 2 ≤ b ≤ 5:

If they are then placed in numerical order, with any repeats removed, we get the following sequence of 15 distinct terms:

4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125

How many distinct terms are in the sequence generated by a^b for 2 ≤ a ≤ 100 and 2 ≤ b ≤ 100?
*/

package main

import "testing"

func TestCountDistinctPowers(t *testing.T) {
	type args struct {
		last_a int64
		last_b int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{last_a: int64(5), last_b: int64(5)}, want: 15},
		{name: "simple", args: args{last_a: int64(100), last_b: int64(100)}, want: 9183},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountDistinctPowers(tt.args.last_a, tt.args.last_b); got != tt.want {
				t.Errorf("CountDistinctPowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
