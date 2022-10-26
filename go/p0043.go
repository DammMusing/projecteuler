/**
Problem 43: Sub-string divisibility

The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the digits 0 to 9 in some order, but it also has a rather interesting sub-string divisibility property.

Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:

    d2d3d4=406 is divisible by 2
    d3d4d5=063 is divisible by 3
    d4d5d6=635 is divisible by 5
    d5d6d7=357 is divisible by 7
    d6d7d8=572 is divisible by 11
    d7d8d9=728 is divisible by 13
    d8d9d10=289 is divisible by 17

Find the sum of all 0 to 9 pandigital numbers with this property.
*/

package main

// Super fast solution (0.17 seconds) by building up pandigitals from a bag of digits, with
// incremental checks for whether the leading string satisfies the divisibility expected.
// Short-circuits any that fail on early factoring checks.
func SumPandigitalsWithFibonacciFactors() int64 {
	var sum int64 = 0

	pandigital10_chan := make(chan int64)
	go pandigital10_sequence(pandigital10_chan)
	for pandigital := <-pandigital10_chan; pandigital != 0; pandigital = <-pandigital10_chan {
		sum += pandigital
	}

	return sum
}

func pan10_helper(seq chan int64, digits []int, avail map[int]bool) {
	factors := []int{2, 3, 5, 7, 11, 13, 17}
	pan10 := int64(0)
	if len(digits) == 9 {
		for di := 0; di < 9; di++ {
			pan10 = (pan10 * 10) + int64(digits[di])
		}
	}
	for i := 0; i < 10; i++ {
		if !avail[i] {
			// only use each digit once
			continue
		}
		if len(digits) >= 3 {
			// check whether the most recently built triplet satisfies the divisibility
			value := digits[len(digits)-2]*100 + digits[len(digits)-1]*10 + i
			if value%factors[len(digits)-3] != 0 {
				continue
			}
		}
		if len(digits) == 9 {
			// if last digit, emit and don't recurse
			seq <- pan10*10 + int64(i)
			continue
		}
		// ...find more digits
		avail[i] = false
		pan10_helper(seq, append(digits, i), avail)
		avail[i] = true
	}
	if len(digits) == 0 {
		// sentinel value for the consumer to stop iterating
		seq <- 0
	}
}

func pandigital10_sequence(seq chan int64) {
	avail := make(map[int]bool)
	for i := 0; i < 10; i++ {
		avail[i] = true
	}
	digits := make([]int, 0)
	pan10_helper(seq, digits, avail)
}
