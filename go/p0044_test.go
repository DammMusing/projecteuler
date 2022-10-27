/**
Problem 44: Pentagon numbers

Pentagonal numbers are generated by the formula, Pn=n(3n−1)/2. The first ten pentagonal numbers are:

        1, 5, 12, 22, 35, 51, 70, 92, 117, 145, ...

It can be seen that P4 + P7 = 22 + 70 = 92 = P8. However, their difference, 70 − 22 = 48, is not pentagonal.

Find the pair of pentagonal numbers, Pj and Pk, for which their sum and difference are pentagonal and D = |Pk − Pj| is minimised; what is the value of D?
*/

package main

import "testing"

func TestClosestPentagonalDistance(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "answer", args: args{limit: 3000}, want: 5482660},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClosestPentagonalDistance(tt.args.limit); got != tt.want {
				t.Errorf("ClosestPentagonalDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
