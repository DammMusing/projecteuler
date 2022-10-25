/**
Problem 42: Coded triangle numbers

The nth term of the sequence of triangle numbers is given by, t_n = ½n(n+1); so the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value. For example, the word value for SKY is 19 + 11 + 25 = 55 = t_10. If the word value is a triangle number then we shall call the word a triangle word.

Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common English words, how many are triangle words?
*/

package main

import (
	"os"
	"strings"
)

func CountCodedTriangleNumbersFromWords() int {
	count := 0
	triangles := make(map[int]bool)
	for i := 0; i < 1000; i++ {
		t_i := (i + 1) * i / 2
		triangles[t_i] = true
	}

	filedata, err := os.ReadFile("../data/p042_words.txt")
	if err != nil {
		panic(err)
	}

	words := strings.Split(string(filedata), `","`)
	words[0] = strings.TrimLeft(words[0], `"`)
	words[len(words)-1] = strings.TrimRight(words[len(words)-1], `"`)

	for _, word := range words {
		if triangles[word_value(word)] {
			count++
		}
	}
	return count
}

func word_value(word string) int {
	value := 0
	for _, ch := range word {
		value += int(ch-'A') + 1
	}

	return value
}
