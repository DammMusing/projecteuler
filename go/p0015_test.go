/**
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

(across across down down)
(across down across down)
(across down down across)
(down across across down)
(down across down across)
(down down across across)

How many such routes are there through a 20×20 grid?
*/

package main

import "testing"

func TestCountRoutesTopLeftToBottomRight(t *testing.T) {
	type args struct {
		size int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "simple", args: args{size: 2}, want: "6"},
		{name: "answer", args: args{size: 20}, want: "137846528820"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountRoutesTopLeftToBottomRight(tt.args.size); got != tt.want {
				t.Errorf("CountRoutesTopLeftToBottomRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
