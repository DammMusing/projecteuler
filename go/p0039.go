/**
Problem 39: Integer right triangles

If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.

{20,48,52}, {24,45,51}, {30,40,50}

For which value of p ≤ 1000, is the number of solutions maximised?
*/

package main

func MostNumerousIntegerRightTrianglePerimeter(limit int) int {
	squares := make(map[int]int)
	for i := 2; i < limit; i++ {
		squares[i*i] = i
	}

	perimeter_counts := make(map[int]int)
	for a := 2; a < limit; a++ {
		for b := a + 1; b < limit; b++ {
			c2 := a*a + b*b
			c := squares[c2]
			if c > 0 {
				perimeter_counts[a+b+c] += 1
			}
		}
	}

	max_perimeter, max_count := 0, 0
	for p, count := range perimeter_counts {
		if p >= limit {
			continue
		}
		if count > max_count {
			max_perimeter = p
			max_count = count
		}
	}
	return max_perimeter
}
