package main

import "testing"

func TestGreet(t *testing.T) {
	expected := "Hello, World!"
	got := Greet()

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
