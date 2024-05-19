package utils

import (
    "fmt"
)

// Deck struct to manage the card deck
type Deck struct {
	cards       []string
	originalDeck []string
}

// NewDeck creates a new Deck with a full set of cards
func NewDeck() *Deck {
	originalDeck := []string{
		"AH", "KH", "QH", "JH", "TH", "9H", "8H", "7H", "6H", "5H", "4H", "3H", "2H",
		"AC", "KC", "QC", "JC", "TC", "9C", "8C", "7C", "6C", "5C", "4C", "3C", "2C",
		"AD", "KD", "QD", "JD", "TD", "9D", "8D", "7D", "6D", "5D", "4D", "3D", "2D",
		"AS", "KS", "QS", "JS", "TS", "9S", "8S", "7S", "6S", "5S", "4S", "3S", "2S",
	}
	return &Deck{
		cards:       append([]string(nil), originalDeck...), // make a copy of the original deck
		originalDeck: originalDeck,
	}
}

// findIndex finds the index of an element in a slice
func (d *Deck) findIndex(element string) int {
	for i, v := range d.cards {
		if v == element {
			return i
		}
	}
	return -1 // return -1 if the element is not found
}

// removeElement removes an element from a slice at a given index
func (d *Deck) removeElement(index int) {
	if index < 0 || index >= len(d.cards) {
		return // do nothing if index is out of range
	}
	d.cards = append(d.cards[:index], d.cards[index+1:]...)
}

// SelectCard selects and removes a card from the deck
func (d *Deck) SelectCard(card string) bool {
	index := d.findIndex(card)
	if index != -1 {
		d.removeElement(index)
		return true
	}
	return false // return false if the card is not found
}

// Reset resets the deck to its original state
func (d *Deck) Reset() {
	d.cards = append([]string(nil), d.originalDeck...) // reset to the original deck
}

func (d *Deck) GetCards() []string {
    return d.cards
}

func (d *Deck) Test() {
    fmt.Println("test")
}
