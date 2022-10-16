/**
Problem 35: Circular primes

The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
*/

package main

import (
	"reflect"
	"testing"
)

func Test_digit_rotations(t *testing.T) {
	type args struct {
		value uint64
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "197", args: args{value: uint64(197)}, want: []int{197, 971, 719}},
		{name: "13542", args: args{value: uint64(13542)}, want: []int{13542, 35421, 54213, 42135, 21354}},
		{name: "2", args: args{value: uint64(2)}, want: []int{2}},
		{name: "11", args: args{value: uint64(11)}, want: []int{11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit_rotations(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("digit_rotations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountCircularPrimesBelow(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{limit: 100}, want: 13},
		{name: "answer", args: args{limit: 1000000}, want: 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountCircularPrimesBelow(tt.args.limit); got != tt.want {
				t.Errorf("CountCircularPrimesBelow() = %v, want %v", got, tt.want)
			}
		})
	}
}
