/**
Problem 31: Coin sums

In the United Kingdom the currency is made up of pound (£) and pence (p). There are eight coins in general circulation:

1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), and £2 (200p).
It is possible to make £2 in the following way:

1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?
*/

package main

func CountWaysToMakeTwoPoundsSterling() int {
	// This includes the single £2 note, all other calculations assume that
	// zero or more 1p coins will make up the difference.
	count := 1

	// Count zero one two ... one hundred two-pence coins.
	for p2 := 0; p2 <= 200; p2 += 2 {
		count += count_5pence(p2)
	}

	return count
}

func count_5pence(subtotal int) int {
	count := 0
	for p5 := 0; subtotal+p5 <= 200; p5 += 5 {
		count += count_10pence(subtotal + p5)
	}
	return count
}

func count_10pence(subtotal int) int {
	count := 0
	for p10 := 0; subtotal+p10 <= 200; p10 += 10 {
		count += count_20pence(subtotal + p10)
	}
	return count
}

func count_20pence(subtotal int) int {
	count := 0
	for p20 := 0; subtotal+p20 <= 200; p20 += 20 {
		count += count_50pence(subtotal + p20)
	}
	return count
}

func count_50pence(subtotal int) int {
	count := 0
	for p50 := 0; subtotal+p50 <= 200; p50 += 50 {
		count += count_pound_notes(subtotal + p50)
	}
	return count
}

func count_pound_notes(subtotal int) int {
	if subtotal > 100 {
		return 1 // just coins
	}
	if subtotal > 0 {
		return 2 // coins or bill + coins
	}
	if subtotal == 0 {
		return 3 // all 1p coins, bill + 100 1p, or 2 bills
	}
	return 0
}
