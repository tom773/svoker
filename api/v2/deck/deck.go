package deck

import (
	"math/rand/v2"
	"strconv"
)

type Card struct {
	Rank string
	Suit string
}

var suits = []string{"H", "D", "C", "S"}

type Deck [52]Card

func NewDeck() Deck {
	d := Deck{}

	cnum := 0
	for j := 0; j < 4; j++ {
		for i := 0; i < 13; i++ {
			d[cnum] = newCard(strconv.Itoa(i+1), suits[j])
			cnum += 1
		}
	}
	return shuffle(d)
}

func newCard(rank string, suit string) Card {
	switch rank {
	case "1":
		rank = "A"
	case "10":
		rank = "T"
	case "11":
		rank = "J"
	case "12":
		rank = "Q"
	case "13":
		rank = "K"
	}
	return Card{rank, suit}
}

func shuffle(d Deck) Deck {

	// Fisher-Yates shuffle
	for i := len(d) - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		d[i], d[j] = d[j], d[i]
	}

	return d
}
