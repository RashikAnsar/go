package main

import "testing"

func TestFactorial(t *testing.T) {
	got := factorial(5)
	want := 120

	if got != want {
		t.Errorf("got %v want %q", got, want)
	}

	got = factorial(0)
	want = 1
	if got != want {
		t.Errorf("got %v want %q", got, want)
	}
}
