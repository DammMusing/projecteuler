/**
Problem 47: Distinct primes factors

The first two consecutive numbers to have two distinct prime factors are:

14 = 2 × 7
15 = 3 × 5

The first three consecutive numbers to have three distinct prime factors are:

644 = 2² × 7 × 23
645 = 3 × 5 × 43
646 = 2 × 17 × 19.

Find the first four consecutive integers to have four distinct prime factors each. What is the first of these numbers?
*/

package main

import (
	"reflect"
	"testing"
)

func TestConsecutiveIntegersWithDistinctPrimes(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "simple", args: args{count: 2}, want: []int{14, 15}},
		{name: "three", args: args{count: 3}, want: []int{644, 645, 646}},
		{name: "answer", args: args{count: 4}, want: []int{134043, 134044, 134045, 134046}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConsecutiveIntegersWithDistinctPrimes(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConsecutiveIntegersWithDistinctPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
