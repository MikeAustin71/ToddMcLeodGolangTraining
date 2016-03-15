package main

import (
	"fmt"
	"math"
)

// More efficient Brute force approach. This will increment
// test number by series maximum.  Speed is acceptable but slows
// down significantly with upper limits >= 30. Note: analysis is
// based on uint64.
//
// See setup info at bottom of this file.

func main() {

	seriesMaxNum := uint64(20)

	series := getSeries(seriesMaxNum)

	smallestDividend, success := findSmallestDividend(series)

	if success {
		fmt.Println("Success. The smallest dividend evenly divisible by the series 1 to ", seriesMaxNum, " is - ", smallestDividend)

	} else {
		fmt.Println("Failure - Did Not Locate smallest dividend.")
	}

}

func getSeries(seriesMaxNum uint64) []uint64 {

	series := make([]uint64, seriesMaxNum-1)

	for i := uint64(2); i <= seriesMaxNum; i++ {
		series[i-2] = i
	}

	return series
}

func findSmallestDividend(series []uint64) (uint64, bool) {

	increment := series[len(series)-1]
	for i := series[len(series)-1]; i < math.MaxUint64; i += increment {

		if IsTestNumDivisbleBySeries(series, i) {
			return i, true
		}
	}

	return 0, false
}

func IsTestNumDivisbleBySeries(series []uint64, testNum uint64) bool {

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
