package client

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tom773/casgo/deck"
)

// Simulating requests from a table to the api
type Player struct {
	Id string
}

type Table struct {
	Id      string
	Deck    deck.Deck
	Players []Player
}

func main() {
	u := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatalf("could not connect to %v: %v", u, err)
	}
	defer conn.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	t := NewTable("table1")

	if err := dealRequest(conn, t); err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("received: %s", message)
		}
	}()
	for {
		select {
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection
			if err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func NewTable(id string) Table {
	return Table{
		Id:   id,
		Deck: deck.Deck{},
		Players: []Player{
			{Id: "player1"},
			{Id: "player2"},
			{Id: "player3"},
		},
	}
}

func dealRequest(conn *websocket.Conn, t Table) error {

	request := map[string]interface{}{
		"type":  "deal",
		"table": t,
	}
	return conn.WriteJSON(request)
}
