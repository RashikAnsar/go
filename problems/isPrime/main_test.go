package main

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	got := isPrime(5)
	want := true

	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}

	got = isPrime(4)
	want = false

	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}
	got = isPrime(97)
	want = true

	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}
}
