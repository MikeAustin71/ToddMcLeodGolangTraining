package main

import (
	"fmt"
	"math/big"
)

// Optimum solution. Very fast and able to process very large
// numbers.
//
// This solution will factor prime number components for series upper limit
// values and compute the smallest dividend evenly divisible by the entire
// series. The analysis uses big.Int numeric types (math/big) and is capable
// of processing very large numbers.
//
// Example: Generate a series with upper limit = 150.
// Smallest Dividend Is:
//  4963595372164418730243844250278933730416682962970482173955824000
//
// See setup info at bottom of this file.

func main() {

	upperLimit := big.NewInt(20)
	series := generateSequence(upperLimit)
	primes := generatePrimes(upperLimit)
	primeExps := calculateMaxExponentsForPrimeSeries(primes, upperLimit)

	fmt.Println("----------------------------------------------")
	fmt.Println("Number series 1 -", upperLimit)

	/*
		fmt.Println("Primes Factors and Prime Exponents --------------------")
		for i := 0; i < len(primes); i++ {
			fmt.Println("Prime: ", primes[i],
				"  Prime Max Exponent: ", primeExps[i])
		}
	*/

	fmt.Println("----------------------------------------------")

	result := computeSmallestDividend(primes, primeExps)

	fmt.Println("Smallest Dividend Is: ", result)

	if IsTestNumDivisibleBySeries(series, result) {
		fmt.Println("Result Confirmed! Result is divisible by each number in the series")
	} else {
		fmt.Print("Failure!!!! - Result is NOT divisible by each number in the series")
	}

}

func generateSequence(upperLimit *big.Int) []*big.Int {

	series := make([]*big.Int, upperLimit.Int64())
	aryLimit := upperLimit.Int64()
	for i := int64(0); i < aryLimit; i++ {

		series[i] = big.NewInt(i + int64(1))
	}

	return series
}

func generatePrimes(upperLimit *big.Int) []*big.Int {
	primes := make([]*big.Int, 0, 100)
	var isPrime bool
	var j int
	aryLimit := upperLimit
	primes = append(primes, big.NewInt(2))
	increment := big.NewInt(2)
	zero := big.NewInt(0)
	modulo := big.NewInt(1)

	for i := big.NewInt(3); i.Cmp(aryLimit) <= 0; i.Add(i, increment) {
		j = 0
		isPrime = true
		y := big.NewInt(1)

		for y = y.Mul(primes[j], primes[j]); y.Cmp(i) <= 0 && j < len(primes); j++ {
			if modulo.Mod(i, primes[j]).Cmp(zero) == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, big.NewInt(0).Set(i))
		}
	}

	return primes
}

func calculateMaxExponentsForPrimeSeries(primes []*big.Int,
	upperLimit *big.Int) []*big.Int {

	primeExps := make([]*big.Int, len(primes))
	limit := len(primes)
	for i := 0; i < limit; i++ {

		primeExps[i] = findLargestExponent(primes[i], upperLimit)

	}

	return primeExps
}

func findLargestExponent(base *big.Int, upperLimit *big.Int) *big.Int {

	exp := big.NewInt(0)
	increment := big.NewInt(1)

	for value := big.NewInt(0).Set(base); value.Cmp(upperLimit) <= 0; exp.Add(exp, increment) {

		value = value.Mul(value, base)

	}

	return exp
}

func computeSmallestDividend(primes []*big.Int,
	primeExponents []*big.Int) *big.Int {

	result := big.NewInt(1)
	limit := len(primes)
	num := big.NewInt(1)

	for i := 0; i < limit; i++ {

		result = result.Mul(result, num.Exp(primes[i], primeExponents[i], nil))
	}

	return result
}

func IsTestNumDivisibleBySeries(series []*big.Int, testNum *big.Int) bool {

	zero := big.NewInt(0)

	for _, v := range series {

		if big.NewInt(0).Mod(testNum, v).Cmp(zero) != 0 {
			return false
		}

	}

	return true
}

/*

* Source: https://projecteuler.net/problem=5

* Title: Smallest multiple - Problem 5

* Problem Description

-----------------------------------------------------------------------
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder. What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
-----------------------------------------------------------------------
Success, Correct Answer = 232792560 - Confirmed by Project Euler

*/
