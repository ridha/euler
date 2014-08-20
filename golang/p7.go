// What is the 10001st prime number?

package main

import (
	"fmt"
	"github.com/ridha/euler/golang/helpers"
)

func main() {
	count := 0
	var i int64 = 3
	for ; ; i += 2 {
		if helpers.IsPrime(i) {
			count += 1
		}
		if count == 10000 {
			fmt.Println(i)
			break
		}
	}
}
