// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143 ?

package main

import (
	"fmt"
	"math"
)

func isPrime(num int64) bool {
	var i int64
	for i = 2; i < int64(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	const num int64 = 600851475143
	var i int64 = int64(math.Sqrt(float64(num)))
	for ; i > 1; i-- {
		if num%i == 0 && isPrime(i) {
			fmt.Println(i)
			break
		}
	}
}
