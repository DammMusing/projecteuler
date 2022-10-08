/**
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

(across across down down)
(across down across down)
(across down down across)
(down across across down)
(down across down across)
(down down across across)

How many such routes are there through a 20×20 grid?
*/

package main

// These numbers get big fast so let's use a BigInt
import "math/big"

// The deeper insight for this problem is recognizing the binomial theorem
// in the routes available.  While it's tempting to think of 'six' in the
// example solution as a selection problem from combinatorics, it's actually
// a constrained selection problem because as soon as you reach the bottom
// wall or right wall it devolves into a single possible path to the end.
//
// The routes for a 2x2 grid are 1 + 2 + 2 + 1
// 1 (follow left wall to bottom)
// 2 (middle via left wall)
// 2 (middle via top wall)
// 1 (follow top to right wall)
//
// The model can be thought of as labeling the intersections of the grid:
// 1 . 1 . 1 . 1 . 1
// .   .   .   .   .
// 1 . 2 . 3 . 4 . 5
// .   .   .   .   .
// 1 . 3 . 6 . 10. 15
// .   .   .   .   .
// 1 . 4 . 10. 20. 35
// .   .   .   .   .
// 1 . 5 . 15. 35. 70
//
// The answer is then the same as asking what the middle of the (n + n + 1)th
// row, or the coefficient for (a + b)^2n when a and b are raised to n.
// e.g. the 6 comes from 6(a^2)(b^2) in (a + b)^4.
//
// The closed form for this is (n Choose n/2) or n! / k! (n-k)! but we can also
// use the big-int Binomial method from golang's standard library.
func CountRoutesTopLeftToBottomRight(size int64) string {
	result := new(big.Int)
	result.Binomial(size*2, size)
	return result.Text(10)
}
