package num

import "math"

/**
 * Implementation of Euler's totient function (Euler's Phi Function)
 * Get all numbers which are coprime to n (starting at 1 to n)
 * @param n
 */
func Totient(n int) []int {
	var coprimes []int
	for i := 1; i <= n; i++ {
		if GCD(n, i) == 1 {
			coprimes = append(coprimes, i)
		}
	}
	return coprimes
}

/**
 * Implementation of the Möbius function.
 * @return
 * 	+1 if n is a square-free positive integer with an even number of prime factors
 * 	-1 if n is a square-free positive integer with an odd number of prime factors
 *	0  if n has a squared prime factor
 */
func Moebius(n int ) int {
	primeFactors := PF(n)
	var primeCount []int
	for i := 0; i < len(primeFactors); i++ {
		v := primeFactors[i]
		if !containsKey(primeCount, v) {
			primeCount = append(primeCount, v)
		} else {
			return 0
		}
	}
	return int(math.Pow(float64(-1), float64(len(primeCount))))
}

/**
 * Calculates the greatest common divisor of two given numbers
 * @param a
 * @param b
 */
func GCD(a int, b int) int {
	r0 := b
	r1 := a % b
	q1 := int(a/b)
	a = q1 * r0 + r1
	if r1 == 0 {
		return r0
	}
	if r1 == 1 {
		return 1
	}
	return GCD(r0, r1)
}

/**
 * Prime factorization of given number n
 * @param n
 * @return Array of the prime factors
 */
func PF(n int) []int {
	var factors []int
	for i := 2; i <= n; i++ {
		for ; n % i == 0;  {
			factors = append(factors, i)
			n /= i
		}
	}
	return factors
}

/**
 * Returns how many times a prime p exists in the factorization of a given number n
 * @param n The number which should be factorized
 * @param p The prime factor in question
 * @return Number of times the prime exists in the factorization of 'n'
 */
func PFExponent(n int, p int) int {
	factors := PF(n)
	counter := 0
	for i := 0; i < len(factors); i++ {
		v := factors[i]
		if p == v {
			counter ++
		}
	}
	return counter
}

/**
 * Check if a key exists in the given array
 * @param array
 * @param n Search for this key
 * @return true if the key exists, false otherwise
 */
func containsKey(array []int, n int) bool {
	for i := 0; i < len(array); i++ {
		if n == array[i] {
			return true
		}
	}
	return false
}
