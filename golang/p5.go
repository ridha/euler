// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import (
	"fmt"
	"github.com/ridha/euler/golang/helpers"
)

func uniprimefact(num int64) map[int64]int {
	factors := make(map[int64]int)

	for _, prime := range helpers.PrimeFactors(int64(num)) {
		factors[prime] = factors[prime] + 1
	}
	return factors
}

func main() {
	num := 20
	factors := make(map[int64]int)

	for i := 2; i <= num; i++ {
		for prime, power := range uniprimefact(int64(i)) {
			if power > factors[prime] {
				factors[prime] = power
			}
		}
	}

	fmt.Println(helpers.Product(factors))
}
