package helpers

import "testing"

func TestIsPrime(t *testing.T) {
	const in, want = 11, true
	if got := IsPrime(in); got != want {
		t.Errorf("IsPrime(%v) = %v, want %v", in, got, want)
	}
}

func TestSqrt(t *testing.T) {
	const in, want = 25, 5
	if got := Sqrt(in); got != want {
		t.Errorf("Sqrt(%v) = %v, want %v", in, got, want)
	}
}
