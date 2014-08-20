package helpers

import "math"

func IsPrime(num int64) bool {
	var i int64
	for i = 2; i <= Sqrt(num); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Sqrt(i int64) int64 {
	return int64(math.Sqrt(float64(i)))
}

//Returns all prime factors of num
func PrimeFactors(num int64) []int64 {
	var (
		i       int64
		factors []int64
	)
	for i = 2; i <= Sqrt(num); {

		if num%i == 0 {
			factors = append(factors, i)
			num = num / i
			continue
		}
		i++
	}

	if num > 1 {
		factors = append(factors, num)
	}

	return factors
}

func Product(nums map[int64]int) uint64 {
	var prod uint64 = 1
	for k, v := range nums {
		prod = prod * uint64(math.Pow(float64(k), float64(v)))
	}
	return prod
}
