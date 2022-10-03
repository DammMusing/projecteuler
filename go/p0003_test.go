/**
Problem 3: Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import "testing"

func TestLargestFactor(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{number: 13195}, want: 29},
		{name: "answer", args: args{number: 600851475143}, want: 6857},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestPrimeFactor(tt.args.number); got != tt.want {
				t.Errorf("LargestFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}
