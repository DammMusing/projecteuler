/**
Problem 48: Self powers

The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.

Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
*/

package main

import "testing"

func TestLastTenDigitsOfSelfPowersTo(t *testing.T) {
	type args struct {
		last int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "example", args: args{last: 10}, want: "405071317"},
		{name: "answer", args: args{last: 1000}, want: "9110846700"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastTenDigitsOfSelfPowersTo(tt.args.last); got != tt.want {
				t.Errorf("LastTenDigitsOfSelfPowersTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
