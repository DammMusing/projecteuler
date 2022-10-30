/**
Problem 49: Prime permutations

The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual in two ways: (i) each of the three terms are prime, and, (ii) each of the 4-digit numbers are permutations of one another.

There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting this property, but there is one other 4-digit increasing sequence.

What 12-digit number do you form by concatenating the three terms in this sequence?
*/

package main

import (
	"testing"
)

func Test_ordered_digits(t *testing.T) {
	type args struct {
		value uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "one", args: args{value: 1487}, want: 1478},
		{name: "two", args: args{value: 4817}, want: 1478},
		{name: "three", args: args{value: 8147}, want: 1478},
		{name: "hmm", args: args{value: 14087}, want: 1478},
		{name: "duplicates", args: args{value: 123123}, want: 112233},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ordered_digits(tt.args.value); got != tt.want {
				t.Errorf("ordered_digits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFourDigitPrimesInArithmeticSequenceWithSameDigits(t *testing.T) {
	tests := []struct {
		name string
		want uint64
	}{
		{name: "answer", want: 296962999629},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FourDigitPrimesInArithmeticSequenceWithSameDigits(); got != tt.want {
				t.Errorf("FourDigitPrimesInArithmeticSequenceWithSameDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
