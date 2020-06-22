package main

import "testing"

func TestExponent(t *testing.T) {
	got := exponent(3, 5)
	want := 243

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
