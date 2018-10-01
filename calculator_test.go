package main

import "testing"

func TestCalculate(t *testing.T) {
	got := event(2)
	expect := 0

	if got != expect {
		t.Error("Exepect event, '%i' but got '%i'", got, expect)
	}
}
