package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/tom773/svoker/api/client"
	"github.com/tom773/svoker/api/deck"
	"github.com/tom773/svoker/api/hub"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var h *hub.Hub

// Comments on this file will be overly verbose because
// I have no idea what the hell is going on with all these
// error as value checks ruining my code flow

func InitWS() {
	h = hub.NewHub()
	go h.Run()

	http.HandleFunc("/ws", wsEp)
	http.HandleFunc("/ws/user", wsuEp)
	http.HandleFunc("/health", healthCheck)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}

// -- End Logic Here, Move to new file once functional

func wsEp(w http.ResponseWriter, r *http.Request) {

	tokenString := r.URL.Query().Get("id")
	if tokenString == "" {
		tokenString = "anon"
	}
	// Accept the WS connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "could not upgrade connection", http.StatusInternalServerError)
	}
	// Hub Stuff For Client Management?
	client_ := &hub.Client{
		ID:   tokenString,
		Conn: conn,
		Send: make(chan []byte, 256),
		Hub:  h,
	}

	h.Register <- client_

	go client_.WritePump()

	// Response Loop Starts
	for {

		// Read an incoming message, close if not.
		_, message, err := conn.ReadMessage()
		if err != nil {
			conn.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
			break
		}
		// Making a map for the request so I can unmarshall it
		request := make(map[string]interface{})

		// Unmarshall the request
		err = json.Unmarshal(message, &request)
		if err != nil {
			log.Println(err)
			continue
		}

		// All WS messages, in and out, go through this switch statement.
		switch request["type"] {
		case "deal":
			// Get the D
			d := deck.NewDeck()
			tableID := request["gameID"].(string)
			hand := client.Deal(d, tableID, h)

			// Marshall the D
			handJSON, err := json.Marshal(hand)
			if err != nil {
				conn.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
				continue
			}
			// Unmarshall the D back into an interface for some reason
			var handsMap map[string]interface{}
			err = json.Unmarshal(handJSON, &handsMap)
			if err != nil {
				conn.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
				continue
			}
			for userID, handData := range handsMap {
				// Marshal the D AGAIN
				handDataJSON, err := json.Marshal(handData)
				if err != nil {
					conn.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
					continue
				}
				h.ToClient(userID, handDataJSON)
			}
		// WS Health Check
		case "health":
			err = conn.WriteMessage(websocket.TextMessage, []byte("ok"))
			if err != nil {
				http.Error(w, "could not write message", http.StatusInternalServerError)
				break
			}
		case "showdown":
			deck := client.Showdown(request["gameID"].(string))
			deckJSON, err := json.Marshal(deck)
			for _, player := range h.Clients {
				h.ToClient(player.ID, deckJSON)
			}
			if err != nil {
				http.Error(w, "could not write message", http.StatusInternalServerError)
				break
			}
		// Reset The Cards
		case "reset":
			client.Reset(request["gameID"].(string))
			reset := map[string]string{
				"reset": "reset",
			}
			resetJSON, err := json.Marshal(reset)
			for _, player := range h.Clients {
				h.ToClient(player.ID, resetJSON)
			}
			if err != nil {
				http.Error(w, "could not write message", http.StatusInternalServerError)
				break
			}
		default:
			err = conn.WriteMessage(websocket.TextMessage, []byte("Unknown Request"))
		}
	}
}
func wsuEp(w http.ResponseWriter, r *http.Request) {
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
			log.Println(err)
		}

		// All WS messages, in and out, go through this switch statement.
		switch request["type"] {
		case "getplayers":
			resp := client.GetPlayers(request["table"].(string))
			respJSON, err := json.Marshal(resp)
			err = conn.WriteMessage(websocket.TextMessage, []byte(respJSON))
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
