/**
Problem 29: Distinct powers

Consider all integer combinations of a^b for 2 ≤ a ≤ 5 and 2 ≤ b ≤ 5:

If they are then placed in numerical order, with any repeats removed, we get the following sequence of 15 distinct terms:

4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125

How many distinct terms are in the sequence generated by a^b for 2 ≤ a ≤ 100 and 2 ≤ b ≤ 100?
*/

package main

import "math/big"

func CountDistinctPowers(last_a int64, last_b int64) int {
	count := 0
	found := make(map[string]bool)
	biglast_a, biglast_b := big.NewInt(last_a), big.NewInt(last_b)
	big_one := big.NewInt(1)

	for a := big.NewInt(2); a.Cmp(biglast_a) <= 0; a.Add(a, big_one) {
		for b := big.NewInt(2); b.Cmp(biglast_b) <= 0; b.Add(b, big_one) {
			x := new(big.Int).Exp(a, b, nil)
			if !found[x.String()] {
				count++
				found[x.String()] = true
			}
		}
	}

	return count
}
