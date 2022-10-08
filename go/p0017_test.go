/**
Problem 17: Number letter counts

If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
*/

package main

import (
	"testing"
)

func TestCountLettersInNumbersUpTo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sssimple", args: args{n: 1}, want: 3},
		{name: "simple", args: args{n: 5}, want: 19},
		{name: "answer", args: args{n: 1000}, want: 21124},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountLettersInNumbersUpTo(tt.args.n); got != tt.want {
				t.Errorf("CountLettersInNumbersUpTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountLettersInNumber(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "ex1", args: args{x: 342}, want: 23},
		{name: "ex2", args: args{x: 115}, want: 20},
		{name: "plain hundred", args: args{x: 100}, want: 10},
		{name: "plain hundred", args: args{x: 300}, want: 12},
		{name: "thousand", args: args{x: 1000}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountLettersInNumber(tt.args.x); got != tt.want {
				t.Errorf("CountLettersInNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
