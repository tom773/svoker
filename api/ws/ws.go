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

func handleMessage(msg map[string]interface{}, game *game.Game, app *pocketbase.PocketBase) {

	response := make(map[string]interface{})
	switch msg["type"] {
	case "deal":
		game.Deck.PrintDeck()
		response := make(map[string]interface{})
		for client := range clients {
			response["type"] = "dealResponse"
			response["cards"] = game.Deck.DrawCard(2)
			game.Action = "flop"
			err := client.WriteJSON(response)
			if err != nil {
				log.Fatal(err)
				client.Close()
				delete(clients, client)
			}
		}
	case "com":
		for client := range clients {
			response["type"] = "comResponse"
			if game.Action == "flop" {
				response["cards"] = game.Deck.DrawCard(3)
				response["action"] = game.Action
				err := client.WriteJSON(response)
				game.Action = "turn"
				if err != nil {
					log.Fatal(err)
					client.Close()
					delete(clients, client)
				}
			} else if game.Action == "turn" {
				response["cards"] = game.Deck.DrawCard(1)
				response["action"] = game.Action
				err := client.WriteJSON(response)
				game.Action = "river"
				if err != nil {
					log.Fatal(err)
					client.Close()
					delete(clients, client)
				}
			} else if game.Action == "river" {
				response["cards"] = game.Deck.DrawCard(1)
				response["action"] = game.Action
				err := client.WriteJSON(response)
				game.Action = "deal"
				game.Deck.Reset()
				if err != nil {
					log.Fatal(err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	default:
		response["error"] = "Invalid message type"
	}
}

func welcomeMessage() map[string]interface{} {
	message := make(map[string]interface{})
	message["type"] = "welcome"
	message["msg"] = "Welcome to the WebSocket server!"
	return message
}
