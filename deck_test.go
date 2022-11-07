package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Error("Expected a deck of length 16, but got:", len(d))
	}
}
