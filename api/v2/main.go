package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/tom773/casgo/client"
	"github.com/tom773/casgo/deck"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Comments on this file will be overly verbose because
// I have no idea what the hell is going on with all these
// error as value checks ruining my code flow

func main() {

	http.HandleFunc("/ws", wsEp)
	http.HandleFunc("/health", healthCheck)
	http.ListenAndServe(":8080", nil)
}

func wsEp(w http.ResponseWriter, r *http.Request) {
	// Accept the WS connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "could not upgrade connection", http.StatusInternalServerError)
	}
	defer conn.Close()

	// Response Loop Starts
	for {

		// Read an incoming message, close if not.
		_, message, err := conn.ReadMessage()
		if err != nil {
			conn.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
			return
		}
		// Making a map for the request so I can unmarshall it
		request := make(map[string]interface{})

		// Unmarshall the request
		err = json.Unmarshal(message, &request)
		if err != nil {
			log.Fatal(err)
		}

		// All WS messages, in and out, go through this switch statement.
		switch request["type"] {

		case "deal":
			// Get the D
			d := deck.NewDeck()
			tableID := request["gameID"].(string)
			client.Deal(d, tableID)
			// Marshall the D to prepare for response
			response, err := json.Marshal(map[string]interface{}{"Event Fired": "Deal"})
			if err != nil {
				conn.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
				continue
			}

			// Write the response to the WS Client
			// Maybe we don't send anything back to the client?
			err = conn.WriteMessage(websocket.TextMessage, response)
			if err != nil {
				http.Error(w, "could not write message", http.StatusInternalServerError)
				break
			}
		// WS Health Check
		case "health":
			err = conn.WriteMessage(websocket.TextMessage, []byte("OK"))
			if err != nil {
				http.Error(w, "could not write message", http.StatusInternalServerError)
				break
			}
		// Reset The Cards
		case "reset":
			client.Reset(request["gameID"].(string))
			err = conn.WriteMessage(websocket.TextMessage, []byte("Resetting"))
			if err != nil {
				http.Error(w, "could not write message", http.StatusInternalServerError)
				break
			}
		default:
			err = conn.WriteMessage(websocket.TextMessage, []byte("Unknown Request"))
		}
	}
}

// Health Check Endpoint for HTTP
func healthCheck(w http.ResponseWriter, r *http.Request) {

	healthResponse := map[string]string{"status": "ok"}
	response, err := json.Marshal(healthResponse)
	if err != nil {
		http.Error(w, "could not marshal health response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}
