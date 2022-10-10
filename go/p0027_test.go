/**
Problem 27: Quadratic primes

Euler discovered the remarkable quadratic formula:

n^2 + n + 41

It turns out that the formula will produce 40 primes for the consecutive integer values . However, when n=40, the result is divisible by 41, and certainly when n=41 it is clearly divisible by 41.

The incredible formula n^2 - 79n + 1601 was discovered, which produces 80 primes for the consecutive values 0 <= n <= 79. The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

   n^2 + an + b, where |a| < 1000 and |b| <= 1000

   where |n| is modulus/absolute value of n
	 e.g. |11| = 11 and |-4| = 4

Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for consecutive values of n, starting with n=0.
*/

package main

import (
	"testing"
)

func TestProductOfCoefficientsForQuadraticPrimeFormula(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{limit: 42}, want: -41},
		{name: "simple", args: args{limit: 1000}, want: -59231},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProductOfCoefficientsForQuadraticPrimeFormula(tt.args.limit); got != tt.want {
				t.Errorf("ProductOfCoefficientsForQuadraticPrimeFormula() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_count_primes_from_quadratic(t *testing.T) {
	type args struct {
		a      int
		b      int
		primes map[int]bool
	}
	primes := get_primes(1000)
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "euler", args: args{a: 1, b: 41, primes: primes}, want: 40},
		{name: "sample", args: args{a: -79, b: 1601, primes: primes}, want: 80},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := count_primes_from_quadratic(tt.args.a, tt.args.b, tt.args.primes); got != tt.want {
				t.Errorf("count_primes_from_quadratic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func get_primes(limit int) map[int]bool {
	primes := make(map[int]bool)
	primechan := make(chan uint64)
	go GatherPrimesLessThan(uint64(2*limit*limit), primechan)
	for prime := <-primechan; prime != 0; prime = <-primechan {
		primes[int(prime)] = true
	}
	return primes
}
