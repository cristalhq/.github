package main

import "testing"

func TestSimple(t *testing.T) {
	if 1 == 0 {
		t.Fatal("wat")
	}
}
