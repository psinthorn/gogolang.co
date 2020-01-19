package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello go"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
