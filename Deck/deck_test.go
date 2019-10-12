package main

import (
	"os"
	"testing"
)

func TestWholeDeck(t *testing.T) {
	d := wholeDeck()

	if len(d) != 52 {
		t.Errorf("Expected length 52, got length %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected Ace of Hearts, got %v", d[0])
	}
}

func TestIO(t *testing.T) {
	os.Remove("_deckTest")

	d := wholeDeck()
	d.savetoFile("_deckTest")

	loadedDeck := readfromFile("_deckTest")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected length 52, got length %v", len(loadedDeck))
	}
}
