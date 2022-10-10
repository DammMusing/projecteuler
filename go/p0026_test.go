/**
Problem 26: Reciprocal cycles

A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

1/2	= 	0.5
1/3	= 	0.(3)
1/4	= 	0.25
1/5	= 	0.2
1/6	= 	0.1(6)
1/7	= 	0.(142857)
1/8	= 	0.125
1/9	= 	0.(1)
1/10	= 	0.1
Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
*/

package main

import (
	"reflect"
	"testing"
)

func Test_digits_of(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "small", args: args{value: 5}, want: []int{5}},
		{name: "large", args: args{value: 1025}, want: []int{1, 0, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digits_of(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("digits_of() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_long_div_pattern(t *testing.T) {
	type args struct {
		divisor int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "trivial", args: args{divisor: 2}, want: nil},
		{name: "trivial2", args: args{divisor: 10}, want: nil},
		{name: "simple", args: args{divisor: 3}, want: []int{3}},
		{name: "eleven", args: args{divisor: 11}, want: []int{0, 9}},
		{name: "seven", args: args{divisor: 7}, want: []int{1, 4, 2, 8, 5, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := long_div_pattern(tt.args.divisor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("long_div_pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pattern_find(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{values: []int{0, 3, 3, 3}}, want: 1},
		{name: "simple2", args: args{values: []int{0, 9, 0, 9, 0, 9}}, want: 2},
		{name: "seven", args: args{values: []int{
			1, 4, 2, 8, 5, 7, 1, 4, 2, 8, 5, 7, 1, 4, 2, 8, 5, 7}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pattern_find(tt.args.values); got != tt.want {
				t.Errorf("pattern_find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestRecurringReciprocalCycle(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{limit: 10}, want: 7},
		{name: "answer", args: args{limit: 1000}, want: 983},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestRecurringReciprocalCycle(tt.args.limit); got != tt.want {
				t.Errorf("LongestRecurringReciprocalCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
