package main

import "testing"

func TestGCD(t *testing.T) {
	got := GCD(36, 60)
	want := 12

	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}

	got = GCD(81, 153)
	want = 9

	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}
}

func TestLCM(t *testing.T) {
	got := LCM(15, 20)
	want := 60

	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}

	got = LCM(72, 120)
	want = 360

	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}
}
