package main

import (
	"testing"
)

func TestSumCalc(t *testing.T) {
	if sumCalc(5) != 15 {
		t.Error("failed")
	} else if sumCalc(3) != 6 {
		t.Error("failed")
	}

}

func TestSumCalc60(t *testing.T) {
	if sumCalcDiv36(5) != 8 {
		t.Error("Failed")
	} else if sumCalcDiv36(15) != 60 {
		t.Errorf("Failed! expected %d, actual %d", 60, sumCalcDiv36(15))
	}

}

func TestIsPrime(t *testing.T) {
	if !isPrime(19) {
		t.Error("Failed, 19 is a prime")
	} else if !isPrime(15485863) {
		t.Error("22,338,618 is a prime")
	} else if isPrime(1) {
		t.Error("Failed 1 is not a prime")
	} else if isPrime(-1) {
		t.Error("Primes can not be negative")
	} else if isPrime(0) {
		t.Error("0 is not a prime")
	}
}
