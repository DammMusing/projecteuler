/**
Problem 18: Maximum path sum I

By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

   3
  7 4
 2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle below:

 	            75
	           95 64
	          17 47 82
	         18 35 87 10
	        20 04 82 47 65
	       19 01 23 75 03 34
	      88 02 77 73 07 63 67
	     99 65 04 28 06 16 70 92
	    41 41 26 56 83 40 80 70 33
	   41 48 72 33 47 32 37 16 94 29
    53 71 44 65 25 43 91 52 97 51 14
	 70 11 33 28 77 73 17 78 39 68 17 57
	91 71 52 38 17 14 91 43 58 50 27 29 48
 63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23

NOTE: As there are only 16384 routes, it is possible to solve this problem by trying every route. However, Problem 67, is the same challenge with a triangle containing one-hundred rows; it cannot be solved by brute force, and requires a clever method! ;o)
*/

package main

// This is a classic example of a dynamic programming problem.  Rather than
// brute-forcing each path from the top down, aggregate maximum values from
// the bottom up, greedily accepting maximum values.  This is because the
// sub-tree's maximum path will indicate which leg of the route the parent
// should traverse, regardless of the value of the immediate child.
//
// It's possible this could be done faster by hand, especially if on paper
// instead of typed here... but there is also increased risk of clerical error
// by typing it here, and in calculation, and typing back the answer.  I'll
// settle on a compromise where the inputs aren't directly parsed from a string,
// but still hardcode the problem statement directly (no simple sample answers).
func MaximumPathSumOnMedTriangle() int {
	best_path := max_subpath(
		[]int{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23},
		[]int{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31})
	best_path = max_subpath(best_path,
		[]int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48})
	best_path = max_subpath(best_path,
		[]int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57})
	best_path = max_subpath(best_path,
		[]int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14})
	best_path = max_subpath(best_path,
		[]int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29})
	best_path = max_subpath(best_path, []int{41, 41, 26, 56, 83, 40, 80, 70, 33})
	best_path = max_subpath(best_path, []int{99, 65, 4, 28, 6, 16, 70, 92})
	best_path = max_subpath(best_path, []int{88, 2, 77, 73, 07, 63, 67})
	best_path = max_subpath(best_path, []int{19, 1, 23, 75, 3, 34})
	best_path = max_subpath(best_path, []int{20, 04, 82, 47, 65})
	best_path = max_subpath(best_path, []int{18, 35, 87, 10})
	best_path = max_subpath(best_path, []int{17, 47, 82})
	best_path = max_subpath(best_path, []int{95, 64})
	if best_path[0] > best_path[1] {
		return best_path[0] + 75
	} else {
		return best_path[1] + 75
	}
}

// catamorphism of successive pairs from `leading` added to each entry of `next`.
func max_subpath(leading []int, next []int) []int {
	if len(leading) != 1+len(next) {
		panic("lengths should be adjacent decreasing from leading to next")
	}
	for i, v := range next {
		if leading[i] > leading[i+1] {
			next[i] = v + leading[i]
		} else {
			next[i] = v + leading[i+1]
		}
	}

	// reuses the backing array of what was defined as the next row
	return next
}
