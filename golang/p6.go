// Find the difference between the sum of the squares of the first one hundred natural numbers and
// the square of the sum

package main

import "fmt"

func main() {
	num := 100
	sum := 0
	sum_of_squares := 0
	for i := 1; i <= num; i++ {
		sum += i
		sum_of_squares += i * i
	}
	fmt.Println(sum*sum - sum_of_squares)
}
