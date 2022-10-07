/**
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
(note: a + b + c = 12 here)

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package main

import "testing"

func TestPythagoreanTripletWithSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{n: 12}, want: 60},
		{name: "answer", args: args{n: 1000}, want: 31875000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PythagoreanTripletWithSum(tt.args.n); got != tt.want {
				t.Errorf("PythagoreanTripletWithSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
