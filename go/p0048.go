/**
Problem 48: Self powers

The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.

Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
*/

package main

import "math/big"

func LastTenDigitsOfSelfPowersTo(last int64) string {
	result := big.NewInt(0)
	mod := big.NewInt(10)
	mod = mod.Exp(mod, mod, nil)

	for i := int64(1); i <= last; i++ {
		next := big.NewInt(i)
		next = next.Exp(next, next, mod)
		result.Add(result, next)
		result.Mod(result, mod)
	}
	return result.String()
}
