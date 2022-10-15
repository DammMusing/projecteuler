/**
Problem 39: Integer right triangles

If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.

{20,48,52}, {24,45,51}, {30,40,50}

For which value of p ≤ 1000, is the number of solutions maximised?
*/

package main

import "testing"

func TestMostNumerousIntegerRightTrianglePerimeter(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{limit: 150}, want: 120},
		{name: "answer", args: args{limit: 1000}, want: 840},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MostNumerousIntegerRightTrianglePerimeter(tt.args.limit); got != tt.want {
				t.Errorf("MostNumerousIntegerRightTrianglePerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
