/**
Problem 5: Smallest multiple

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import "testing"

func TestSmallestMultiple(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{limit: 10}, want: 2520},
		{name: "answer", args: args{limit: 20}, want: 232792560},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallestMultiple(tt.args.limit); got != tt.want {
				t.Errorf("SmallestMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}
