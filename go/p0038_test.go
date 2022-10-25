/**
Problem 38: Pandigital multiples

Take the number 192 and multiply it by each of 1, 2, and 3:

192 × 1 = 192
192 × 2 = 384
192 × 3 = 576

By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?
*/

package main

import "testing"

func TestLargestPandigitalFromConcatenatedSerialProduct(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{limit: 20}, want: 918273645},
		{name: "answer", args: args{limit: 10000}, want: 932718654},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestPandigitalFromConcatenatedSerialProduct(tt.args.limit); got != tt.want {
				t.Errorf("LargestPandigitalFromConcatenatedSerialProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
