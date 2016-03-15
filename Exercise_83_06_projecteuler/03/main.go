package main

import (
	"fmt"
	"math"
)

// More efficient approach.
//
// This solution will factor prime number components for series upper limit
// and compute the smallest dividend evenly divisible by the entire series.
//
// This solution is based on uint64 types. Accuracy above upper limit = 40 becomes
// unreliable as result approaches and surpasses the bounds of uint64.
//
// Max uint64 = 18446744073709551615 (20-digits).
//
// See setup info at bottom of this file.

func main() {

	seriesUpperLimit := uint64(20)

	primes := generatePrimes(seriesUpperLimit)

	smallestDividend := findSmallestDividend(primes, seriesUpperLimit)

	series := getSeries(seriesUpperLimit)

	if isTestNumDivisbleBySeries(series, smallestDividend) {
		fmt.Println("The smallest dividend evenly divisible by the series 1 to ", seriesUpperLimit, " is - ", smallestDividend)

	} else {
		fmt.Println("Error!!! Could Not Compute the smallest dividend evenly divisible by the series 1 to ", seriesUpperLimit)
		fmt.Println("Incorrect Result = ", smallestDividend)
	}

}

func getSeries(upperLimit uint64) []uint64 {

	series := make([]uint64, upperLimit-1)

	for i := uint64(2); i <= upperLimit; i++ {
		series[i-2] = i
	}

	return series
}

func findSmallestDividend(primes []uint64, upperLimit uint64) uint64 {

	result := uint64(1)

	for i := 0; i < len(primes); i++ {
		a := math.Floor(
			math.Log(float64(upperLimit)) / math.Log(float64(primes[i])))

		result = result * uint64(math.Pow(float64(primes[i]), float64(a)))
	}

	return result
}

func generatePrimes(upperLimit uint64) []uint64 {
	var primes []uint64
	var isPrime bool
	var j int
	primes = append(primes, 2)

	for i := uint64(3); i <= upperLimit; i += 2 {
		j = 0
		isPrime = true
		for y := primes[j] * primes[j]; y <= i && j < len(primes); j++ {
			if i%primes[j] == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}

func isTestNumDivisbleBySeries(series []uint64, testNum uint64) bool {

	for _, v := range series {

		if testNum%v != 0 {
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
