/**
Problem 40: Champernowne's constant

An irrational decimal fraction is created by concatenating the positive integers:

    0.123456789101112131415161718192021...
		             ^
                 12th digit

It can be seen that the 12th digit of the fractional part is 1.

If d_n represents the nth digit of the fractional part, find the value of the following expression.

d_1 × d_10 × d_100 × d_1000 × d_10000 × d_100000 × d_1000000
*/

package main

func ChampernowneConstant(lg_digits int) int {
	digit_product := 1

	num_digits := 10
	for i := 2; i <= lg_digits; i++ {
		num_digits *= 10
	}
	digit_chan := make(chan int)
	go generate_digits(num_digits, digit_chan)

	i := 0
	for d := range digit_chan {
		i += 1
		for j, k := 2, 10; j <= lg_digits; j, k = j+1, k*10 {
			if i == k {
				digit_product *= d
				break
			}
		}
	}

	return digit_product
}

func generate_digits(limit int, digits chan int) {
	i, count_digits := 1, 1
	for count_digits <= limit {
		for _, d := range digits_of(i) {
			digits <- d
			count_digits++
		}
		i++
	}
	close(digits)
}
