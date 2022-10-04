/**
The sum of the squares of the first ten natural numbers is,

   1^2 + 2^2 + 3^2 ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,

   (1 + 2 + 3 ... + 10)^2 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 - 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

package main

import (
	"math/big"
	"reflect"
	"testing"
)

func TestSumSquareDifference(t *testing.T) {
	type args struct {
		limit int64
	}
	tests := []struct {
		name string
		args args
		want big.Int
	}{
		{name: "simple", args: args{limit: 10}, want: *big.NewInt(2640)},
		{name: "answer", args: args{limit: 100}, want: *big.NewInt(25164150)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumSquareDifference(tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumSquareDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
