/**
Problem 17: Number letter counts

If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
*/

package main

// note: I personally don't like introducing "and" unless a fraction follows the term, but I'm also not British.

func CountLettersInNumbersUpTo(n int) int {
	if n <= 0 {
		panic("only expecting a natural number as limit")
	}
	if n > 1000 {
		panic("expected a limit no greater than one thousand")
	}
	total_sum := 0

	for i := 1; i <= n; i++ {
		total_sum += CountLettersInNumber(i)
	}

	return total_sum
}

func CountLettersInNumber(x int) int {
	if x <= 0 {
		panic("only expecting a natural number")
	}
	if x == 1000 {
		return len("one") + len("thousand")
	}

	digit_length := [20]int{
		0,
		len("one"),
		len("two"),
		len("three"),
		len("four"),
		len("five"),
		len("six"),
		len("seven"),
		len("eight"),
		len("nine"),
		len("ten"),
		len("eleven"),
		len("twelve"),
		len("thirteen"),
		len("fourteen"),
		len("fifteen"),
		len("sixteen"),
		len("seventeen"),
		len("eighteen"),
		len("nineteen"),
	}
	tens_length := [10]int{
		0, 0,
		len("twenty"),
		len("thirty"),
		len("forty"),
		len("fifty"),
		len("sixty"),
		len("seventy"),
		len("eighty"),
		len("ninety"),
	}

	if x < 20 {
		return digit_length[x]
	}

	if x < 100 {
		return tens_length[x/10] + digit_length[x%10]
	}

	if x < 1000 {
		h_sum := digit_length[x/100] + len("hundred")
		if x%100 != 0 {
			h_sum += len("and")
		}
		if x%100 < 20 {
			return h_sum + digit_length[x%20]
		} else {
			return h_sum + tens_length[(x/10)%10] + digit_length[x%10]
		}
	}

	panic("expected a value no greater than one thousand")
}
