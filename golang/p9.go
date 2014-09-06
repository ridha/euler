// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
// a ** 2 + b ** 2 = c ** 2

// For example, 32 + 42 = 9 + 16 = 25 = 52.

// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

package main

import "fmt"

func main() {
	var a, b uint64

	for a = 1; a < 1000; a++ {
		for b = a + 1; b < 1000; b++ {

			c := 1000 - a - b

			if c < 0 {
				continue
			}

			if a*a+b*b == c*c {
				fmt.Println(c * a * b)
			}
		}
	}

}
