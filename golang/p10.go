// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.
// 142913828922

package main

import (
	"fmt"
	"github.com/ridha/euler/golang/helpers"
)

func generate_prime(start int64, end int64, c chan int64) {
	for i := start; i < end; i += 2 {
		if helpers.IsPrime(i) {
			c <- i
		}
	}
	close(c)
}

func main() {
	var sum int64 = 2
	c := make(chan int64, 2)

	go generate_prime(3, 2000000, c)
	for p := range c {
		sum += p
	}

	fmt.Println(sum)
}
