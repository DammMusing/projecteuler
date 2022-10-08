package utils

/**
 * Send all the prime values that exist below `limit` through the provided channel.
 */
func GatherPrimesLessThan(limit uint64, primes chan uint64) {
	// Initialize the sieve, but handle the first prime as part of initialization.
	primes <- 2
	sieve := NewSieve(limit)
	sieve.Set(2, true)
	for i := uint64(3); i < limit; i += 2 {
		sieve.Set(i, true)
	}

	for i := uint64(3); i < limit; i++ {
		if sieve.Get(i) {
			primes <- i // Send the prime along the channel...
			for multiple := i * 2; multiple <= limit; multiple += i {
				// ...and clear all the multiples from the sieve
				sieve.Set(multiple, false)
			}
		}
	}

	// Emit a sentinel value to indicate completion.
	primes <- 0
}

/** A bitmap datatype for representing the Sieve of Eratosthenes. */
type Sieve struct {
	// TODO: this could be made much more efficient by eliding 4 of 6 bits,
	// implicitly avoiding all products of 2 and of 3.  Currently it represents
	// all natural numbers, even zero and one.
	bits []byte
}

func NewSieve(size uint64) *Sieve {
	array_size := size >> 3
	if size&0x07 > 0 {
		array_size++
	}
	return &Sieve{make([]byte, array_size)}
}

func (s *Sieve) Set(index uint64, value bool) {
	if index >= uint64(len(s.bits)<<3) {
		return
	}
	pos, bit := index>>3, byte(1<<(index&0x07))
	if value {
		s.bits[pos] |= bit
	} else {
		s.bits[pos] &= ^bit
	}
}

func (s *Sieve) Get(index uint64) bool {
	if index >= uint64(len(s.bits))<<3 {
		return false
	}
	pos, bit := index>>3, byte(1<<(index&0x07))
	return s.bits[pos]&bit != 0
}

func NumFactorsOf(n uint64) int {
	num_factors := 0

	for x := uint64(1); x < n; x++ {
		if n%x == 0 {
			num_factors += 1
		}
	}

	return num_factors
}
