// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143 ?

package main

import (
	"./helpers"
	"fmt"
)

func main() {
	const num int64 = 600851475143
	var i int64 = helpers.Sqrt(num)
	for ; i > 1; i-- {
		if num%i == 0 && helpers.IsPrime(i) {
			fmt.Println(i)
			break
		}
	}
}
