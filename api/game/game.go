package game

import "github.com/tom773/svoker/api/utils"

type Game struct {
	Deck    *utils.Deck
	TableID string
	Action  string
}

func GameState() *Game {
	game := &Game{
		Deck:    utils.NewDeck(),
		TableID: "1",
		Action:  "deal",
	}
	game.Deck.Reset()
	return game
}
