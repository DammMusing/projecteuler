/**
Problem 35: Circular primes

The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
*/

package main

import (
	"fmt"
	"math"
)

func CountCircularPrimesBelow(limit int) int {
	count := 0

	primes := make(map[int]bool)
	prime_chan := make(chan uint64)
	go GatherPrimesLessThan(uint64(limit), prime_chan)
	for p := <-prime_chan; p != 0; p = <-prime_chan {
		fmt.Printf("prime %d\n", p)
		primes[int(p)] = true
	}
	fmt.Println("Gathered primes, checking for cycles.")

	prime_chan = make(chan uint64)
	go GatherPrimesLessThan(uint64(limit), prime_chan)
	for p := <-prime_chan; p != 0; p = <-prime_chan {
		rotations := digit_rotations(p)
		circular := true
		for _, v := range rotations {
			if !primes[v] {
				circular = false
			}
		}
		if circular {
			fmt.Printf("Found cycle (%d), working rotation: ", p)
			for _, v := range rotations {
				// Reset all primes in the ring so we don't revisit them later
				primes[v] = false
				count += 1
				fmt.Printf("%d ", v)
			}
			fmt.Printf("\n")
		}
	}

	return count
}

func digit_rotations(value uint64) []int {
	if value > math.MaxInt {
		panic("digit_rotations() expects int-sized uint64. deal with it.")
	}
	iv := int(value)
	digits := digits_of(iv)
	rotations := make([]int, 1)
	rotations[0] = iv
	for i := 0; i < len(digits); i++ {
		rotated := from_digits(append(digits[i:], digits[:i]...))
		if rotated == int(value) {
			continue
		}
		rotations = append(rotations, rotated)
	}

	return rotations
}

func from_digits(digits []int) int {
	result := 0
	for i := 0; i < len(digits); i++ {
		result = 10*result + digits[i]
	}
	return result
}
