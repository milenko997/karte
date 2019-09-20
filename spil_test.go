package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v", len(d))
	}
	if d[0] != "Kec Karo"{
		t.Errorf("Expected first card is Kec Karo, but we got %v", d[0])
	}
	if d[len(d)-1] != "Cetvorka Tref"{
		t.Errorf("Expected last card Cetvorka Tref, but we got %v", d[len(d)-1])
	}
}
