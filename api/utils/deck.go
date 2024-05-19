package utils

import (
	"fmt"
	"math/rand/v2"
)

// Deck struct to manage the card deck
type Deck struct {
	Cards        []string
	OriginalDeck []string
}

// NewDeck creates a new Deck with a full set of Cards
func NewDeck() *Deck {
	OriginalDeck := []string{
		"AH", "KH", "QH", "JH", "TH", "9H", "8H", "7H", "6H", "5H", "4H", "3H", "2H",
		"AC", "KC", "QC", "JC", "TC", "9C", "8C", "7C", "6C", "5C", "4C", "3C", "2C",
		"AD", "KD", "QD", "JD", "TD", "9D", "8D", "7D", "6D", "5D", "4D", "3D", "2D",
		"AS", "KS", "QS", "JS", "TS", "9S", "8S", "7S", "6S", "5S", "4S", "3S", "2S",
	}
	return &Deck{
		Cards:        append([]string(nil), OriginalDeck...), // make a copy of the original deck
		OriginalDeck: OriginalDeck,
	}
}

// findIndex finds the index of an element in a slice
func (d *Deck) findIndex(element string) int {
	for i, v := range d.Cards {
		if v == element {
			return i
		}
	}
	return -1 // return -1 if the element is not found
}

// removeElement removes an element from a slice at a given index
func (d *Deck) removeElement(index int) {
	if index < 0 || index >= len(d.Cards) {
		return // do nothing if index is out of range
	}
	d.Cards = append(d.Cards[:index], d.Cards[index+1:]...)
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

// Reset resets the deck to its Original state
func (d *Deck) Reset() {
	d.Cards = append([]string(nil), d.OriginalDeck...) // reset to the original deck
}

func (d *Deck) DrawCard(n int) []string {
	var drawnCards []string
	for i := 0; i < n; i++ {
		if len(d.Cards) == 0 {
			break
		}
		index := rand.IntN(len(d.Cards))
		card := d.Cards[index]
		drawnCards = append(drawnCards, card)
		d.removeElement(index)
	}
	return drawnCards
}

// PrintDeck prints the current deck of cards
func (d *Deck) PrintDeck() {
	fmt.Println(d.Cards)
}
