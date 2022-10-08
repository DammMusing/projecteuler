/**
Problem 22: Names scores

Using ../data/p022_names.txt, a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

package main

import "testing"

func TestTotalNamesScores(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "answer", want: 871198282},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalNamesScores(); got != tt.want {
				t.Errorf("TotalNamesScores() = %v, want %v", got, tt.want)
			}
		})
	}
}
