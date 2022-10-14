/**
Problem 33: Digit cancelling fractions

The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
*/

package main

import "fmt"

func LowestCommonDenominatorOfDigitCancellingFractions() int {
	nums, dens := make([]int, 0), make([]int, 0)

	for num1 := 1; num1 < 10; num1++ {
		for num2 := 1; num2 < 10; num2++ {
			num := num1*10 + num2
			for den1 := num1; den1 < 10; den1++ {
				for den2 := 1; den2 < 10; den2++ {
					den := den1*10 + den2
					if num >= den {
						// must be less than 1 and non-trivial
						continue
					}
					if num1 == den2 {
						before := float32(num) / float32(den)
						after := float32(num2) / float32(den1)
						if before == after {
							fmt.Printf("%d / %d == %d / %d\n", num, den, num2, den1)
							nums = append(nums, num2)
							dens = append(dens, den1)
						}
					}
					if num2 == den1 {
						before := float32(num) / float32(den)
						after := float32(num1) / float32(den2)
						if before == after {
							fmt.Printf("%d / %d == %d / %d\n", num, den, num1, den2)
							nums = append(nums, num1)
							dens = append(dens, den2)
						}
					}
				}
			}
		}
	}

	prod_num, prod_den := 1, 1
	for _, v := range nums {
		prod_num *= v
	}
	for _, v := range dens {
		prod_den *= v
	}

	// fortunately, the denominator is evenly dividable by the numerator.
	return int(prod_den / prod_num)
}
