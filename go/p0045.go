/**
Problem 45: Triangular, pentagonal, and hexagonal

Triangle, pentagonal, and hexagonal numbers are generated by the following formulae:

Triangle	 	Tn=n(n+1)/2	 	1, 3, 6, 10, 15, ...
Pentagonal	 	Pn=n(3n−1)/2	 	1, 5, 12, 22, 35, ...
Hexagonal	 	Hn=n(2n−1)	 	1, 6, 15, 28, 45, ...
It can be verified that T285 = P165 = H143 = 40755.

Find the next triangle number that is also pentagonal and hexagonal.
*/

package main

func TriPentaHexagonNumber(limit int64) int64 {
	result := int64(0)

	hexas := make(map[int64]bool)
	pentas := make(map[int64]bool)

	i, f_i := int64(1), int64(1)
	for f_i < limit {
		i += 1
		f_i = pentagonal64(i)
		pentas[f_i] = true
	}
	i, f_i = 1, 1
	for f_i < limit {
		i += 1
		f_i = hexagonal(i)
		hexas[f_i] = true
	}

	i, f_i = 1, 1
	for f_i < limit {
		i += 1
		f_i = triangular(i)
		if pentas[f_i] && hexas[f_i] {
			result = f_i
		}
	}

	return result
}

func pentagonal64(n int64) int64 {
	return n * (3*n - 1) / 2
}

func hexagonal(n int64) int64 {
	return n * (2*n - 1)
}

func triangular(n int64) int64 {
	return n * (n + 1) / 2
}
