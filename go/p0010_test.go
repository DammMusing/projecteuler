/**
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import "testing"

func TestSumOfPrimesBelow(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "simple", args: args{n: 10}, want: 17},
		{name: "answer", args: args{n: 2000000}, want: 142913828922},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfPrimesBelow(tt.args.n); got != tt.want {
				t.Errorf("SumOfPrimesBelow() = %v, want %v", got, tt.want)
			}
		})
	}
}
