package weby

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pocketbase/pocketbase"
	"github.com/tom773/svoker/api/game"
)

// WebSockets Logic Begins Here
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan map[string]interface{})
var gameStateTest = make(map[string]interface{})

// Exported function to handle incoming connections
func HandleConnection(w http.ResponseWriter, r *http.Request, game *game.Game) {
	app := pocketbase.New()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	clients[conn] = true

	for client := range clients {
		if err := client.WriteJSON(welcomeMessage()); err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(clients, client)
			return
		}
	}

	for {
		var msg map[string]interface{}
		err := conn.ReadJSON(&msg)

		if err != nil {
			log.Fatal(err)
			return
		}

		handleMessage(msg, game, app)
	}
}

type GameStateRes struct {
	Cards  []string `json:"cards"`
	Action string   `json:"action"`
	Type   string   `json:"type"`
	Error  string   `json:"error"`
}

func handleMessage(msg map[string]interface{}, game *game.Game, app *pocketbase.PocketBase) {

	response := GameStateRes{}
	switch msg["type"] {
	case "com":
		var numCards int
		switch game.Action {
		case "preflop":
			response.Type = "dealResponse"
			for client := range clients {
				response.Cards = game.Deck.DrawCard(2)
				err := client.WriteJSON(response)
				if err != nil {
					log.Fatal(err)
					client.Close()
					delete(clients, client)
				}
			}
			game.Action = "flop"
		case "flop":
			numCards = 3
			game.Action = "turn"
		case "turn":
			numCards = 1
			game.Action = "river"
		case "river":
			numCards = 1
			game.Action = "showdown"
		default:
			response.Error = "Mayday, we've had an oopsy woopsy"
			for client := range clients {
				err := client.WriteJSON(response)
				if err != nil {
					log.Fatal(err)
					client.Close()
					delete(clients, client)
				}
			}
			return
		}
		// End of switch statement

		if numCards > 0 {
			response.Cards = game.Deck.DrawCard(numCards)
		}

		response.Type = "comResponse"
		response.Action = game.Action

		for client := range clients {
			err := client.WriteJSON(response)
			if err != nil {
				log.Fatal(err)
				client.Close()
				delete(clients, client)
			}
		}

		if game.Action == "showdown" {
			game.Deck.Reset()
			game.Action = "preflop"
		}
	default:
		response.Error = "Invalid message type"
		for client := range clients {
			err := client.WriteJSON(response)
			if err != nil {
				log.Fatal(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func welcomeMessage() map[string]interface{} {
	message := make(map[string]interface{})
	message["type"] = "welcome"
	message["msg"] = "Welcome to the WebSocket server!"
	return message
}
