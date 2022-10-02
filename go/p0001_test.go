/**
Problem 1: Multiples of 3 or 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import "testing"

func TestDoSum(t *testing.T) {
	type args struct {
		upper_bound int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sanity check", args: args{upper_bound: 10}, want: 23},
		{name: "answer", args: args{upper_bound: 1000}, want: 233168},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoSum(tt.args.upper_bound); got != tt.want {
				t.Errorf("DoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
