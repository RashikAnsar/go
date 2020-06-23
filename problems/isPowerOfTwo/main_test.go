package main

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	got := isPowerOfTwo(32)
	want := true

	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}

	got = isPowerOfTwo(510)
	want = false

	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}

	got = isPowerOfTwo(1024)
	want = true

	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}

}
