package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/tom773/casgo/deck"
)

func TestWSRandomness(t *testing.T) {
	// Shuffle some decks and check if the order is different
	// We'd expect a uniform distribution of the number of times each card appears
	// Easily done within a margin of 10%

	const iterations = 10000
	cardCounts := make(map[string]int)
	// Meat and potatoes of the WS connection
	u := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("could not connect to %v: %v", u, err)
	}
	defer conn.Close()

	for i := 0; i < iterations; i++ {
		request := map[string]interface{}{
			"type": "deal",
		}
		err = conn.WriteJSON(request)
		if err != nil {
			t.Fatal(err)
		}

		_, message, err := conn.ReadMessage()
		if err != nil {
			t.Fatal(err)
		}
		var responseDeck deck.Deck
		err = json.Unmarshal(message, &responseDeck)
		if err != nil {
			t.Fatal(err)
		}
		for _, card := range responseDeck {
			cardCounts[card.Rank+card.Suit]++
		}

	}

	expected := iterations * len(deck.NewDeck()) / 52
	margin := expected / 10

	for card, count := range cardCounts {
		if count < expected-margin || count > expected+margin {
			t.Errorf("card %v appeared %v times, expected %v", card, count, expected)
		}
	}
}

// WS Testing

func TestWSHealth(t *testing.T) {

	u := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("could not connect to %v: %v", u, err)
	}
	defer conn.Close()

	// Sample Data - Will Be Generating Thousands of These In Future

	request := map[string]interface{}{
		"type": "health",
	}

	err = conn.WriteJSON(request)
	if err != nil {
		t.Fatal(err)
	}

	_, message, err := conn.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Got: %s", message)

	var responseDeck deck.Deck

	if len(responseDeck) == 0 {
		t.Error("You suck")
	}
}

func TestHTTPHealth(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/health")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status code 200, got %v", resp.StatusCode)
	}
}
