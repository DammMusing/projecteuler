/**
Problem 49: Prime permutations

The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual in two ways: (i) each of the three terms are prime, and, (ii) each of the 4-digit numbers are permutations of one another.

There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting this property, but there is one other 4-digit increasing sequence.

What 12-digit number do you form by concatenating the three terms in this sequence?
*/

package main

func FourDigitPrimesInArithmeticSequenceWithSameDigits() uint64 {
	prime_chan := make(chan uint64)
	go GatherPrimesLessThan(10000, prime_chan)

	clusters := make(map[int][]int, 0)

	for prime := <-prime_chan; prime != 0; prime = <-prime_chan {
		hash := int(ordered_digits(prime))
		if len(digits_of(hash)) < 4 {
			// For this problem we exclude any 3-digit results.
			continue
		}
		clusters[hash] = append(clusters[hash], int(prime))
	}

	for _, cluster := range clusters {
		if len(cluster) >= 3 {
			triple := has_arithmetic_triple(cluster)
			if triple != nil && triple[0] != 1487 {
				return uint64(triple[0])*uint64(100000000) + uint64(triple[1]*10000) + uint64(triple[2])
			}
		}
	}

	return 0
}

func ordered_digits(value uint64) uint64 {
	digits := make(map[int]int)
	for value > 0 {
		digits[int(value%10)] += 1
		value /= 10
	}
	result := uint64(0)
	for i := uint64(1); i < 10; i++ {
		for j := digits[int(i)]; j > 0; j-- {
			result = result*10 + i
		}
	}
	return result
}

func has_arithmetic_triple(cluster []int) []int {
	for i := 0; i < len(cluster)-2; i++ {
		for j := i + 1; j < len(cluster)-1; j++ {
			for k := j + 1; k < len(cluster); k++ {
				if cluster[j]-cluster[i] == cluster[k]-cluster[j] {
					return []int{cluster[i], cluster[j], cluster[k]}
				}
			}
		}
	}
	return nil
}
