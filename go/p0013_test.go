/**
Work out the first ten digits of the sum of the one-hundred 50-digit numbers.
*/

package main

import "testing"

func TestSumTheseReallyLargeNumbersFirstTenDigits(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "answer", want: "5537376230"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumTheseReallyLargeNumbersFirstTenDigits(); got != tt.want {
				t.Errorf("SumTheseReallyLargeNumbersFirstTenDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
