/**
Problem 22: Names scores

Using ../data/p022_names.txt, a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

package main

import (
	"os"
	"sort"
	"strings"
)

func TotalNamesScores() int {
	total := 0
	filedata, err := os.ReadFile("../data/p022_names.txt")
	check(err)

	names := strings.Split(string(filedata), `","`)
	// trim initial and trailing quotes
	names[0] = strings.TrimLeft(names[0], `"`)
	names[len(names)-1] = strings.TrimRight(names[len(names)-1], `"`)

	// sort names alphabetically
	sort.Strings(names)

	for i, name := range names {
		total += name_score(i+1, name)
	}

	return total
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func name_score(position int, name string) int {
	char_score := 0
	offset := int('A') - 1
	for _, ch := range name {
		char_score += int(ch) - offset
	}

	return char_score * position
}
