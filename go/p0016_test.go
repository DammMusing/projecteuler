/**
Problem 16: Power digit sum

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

package main

import "testing"

func TestSumOfDigitsOfPowerOfTwo(t *testing.T) {
	type args struct {
		power int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{power: 15}, want: 26},
		{name: "answer", args: args{power: 1000}, want: 1366},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfDigitsOfPowerOfTwo(tt.args.power); got != tt.want {
				t.Errorf("SumOfDigitsOfPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
