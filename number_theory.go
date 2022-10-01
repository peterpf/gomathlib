package gomathlib

import "math"

// Totient implements Euler's totient function (Euler's Phi Function) to find all numbers which are coprime to n (starting at 1 to n).
func Totient(n int) []int {
	var coprimes []int
	for i := 1; i <= n; i++ {
		if GCD(n, i) == 1 {
			coprimes = append(coprimes, i)
		}
	}
	return coprimes
}

// Implementation of the Möbius function. Returns
// +1 if n is a square-free positive integer with an even number of prime factors,
// -1 if n is a square-free positive integer with an odd number of prime factors,
// 0 if n has a squared prime factor.
func Moebius(n int) int {
	primeFactors := PrimeFactorization(n)
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

// GCD calculates the greatest common divisor of two given numbers
func GCD(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// PrimeFactorization return an array of the prime factors of the number n.
func PrimeFactorization(n int) []int {
	var factors []int
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	return factors
}

// PrimeFactorExponent calculates how many times a prime `p` exists in the factorization of a given number `n`.
func PrimeFactorExponent(n int, p int) int {
	factors := PrimeFactorization(n)
	counter := 0
	for i := 0; i < len(factors); i++ {
		v := factors[i]
		if p == v {
			counter++
		}
	}
	return counter
}

// containsKey returns true if a key exists in the given array.
func containsKey(array []int, n int) bool {
	for i := 0; i < len(array); i++ {
		if n == array[i] {
			return true
		}
	}
	return false
}
