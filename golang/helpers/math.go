package helpers

import "math"

func IsPrime(num int64) bool {
	var i int64
	for i = 2; i < Sqrt(num); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Sqrt(i int64) int64 {
	return int64(math.Sqrt(float64(i)))
}
