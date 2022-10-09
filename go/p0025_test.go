/**
Problem 25: 1000 digit Fibonacci number

The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
*/

package main

import "testing"

func TestFindFirstFibonacciOfNDigits(t *testing.T) {
	type args struct {
		ndigits int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{ndigits: 3}, want: 12},
		{name: "answer", args: args{ndigits: 1000}, want: 4782},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindFirstFibonacciOfNDigits(tt.args.ndigits); got != tt.want {
				t.Errorf("FindFirstFibonacciOfNDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
