/**
Problem 14: Longest Collatz sequence

The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

package main

func LongestCollatzSequenceBelow(n int) int {
	stack := make([]int, 0, 128) // stack starts at medium capacity but may grow
	lengths := make([]int, n)
	lengths[1] = 1

	for candidate := 2; candidate < n; candidate++ {
		if lengths[candidate] != 0 {
			// we already solved the sequence for this one, goto next candidate
			continue
		}
		// track all the new elements found, for later iterating in reverse (at
		// the time of writing this I didn't know about golang-collections/Stack)
		stack = append(stack, candidate)
		steps := 0
		for {
			candidate = next(candidate)
			if candidate < n && lengths[candidate] != 0 {
				steps = lengths[candidate] + 1
				break
			}
			stack = append(stack, candidate)
			steps++
		}

		// track the step length for each element found in the sequence before
		for i := len(stack) - 1; i > 0; i-- {
			candidate = stack[i]
			if candidate < n {
				lengths[candidate] = steps
			}
			stack = stack[:i]
			steps++
		}
	}

	// All sequence lengths found, search for longest of them
	longest, value := 0, 0
	for i, v := range lengths {
		if v > longest {
			longest, value = v, i
		}
	}
	return value
}

func next(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return (3 * n) + 1
}
