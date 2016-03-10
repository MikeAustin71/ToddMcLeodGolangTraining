package main

import (
	"fmt"
	"math"
)

func main() {
	series := GetSeries()

	smallestDividend, success := FindSmallestDividend(series)

	if success {
		fmt.Println("Success. The smallest dividend evenly divisible by the series 1 to 20 is - ", smallestDividend)

	} else {
		fmt.Println("Failure - Did Not Locate smallest dividend.")
	}

}

func GetSeries() []uint64 {

	var MAXSERIESNUM uint64 = 20

	series := make([]uint64, MAXSERIESNUM-1)

	for i := uint64(2); i <= MAXSERIESNUM; i++ {
		series[i-2] = i
	}

	return series
}

func FindSmallestDividend(series []uint64) (uint64, bool) {

	for i := series[len(series)-1] + 1; i < math.MaxUint64; i++ {

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
