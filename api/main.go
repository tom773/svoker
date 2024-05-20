package main

import (
	"github.com/tom773/svoker/api/routes"
)

func main() {
	routes.SetupRoutes()
}

// This API is now functional for a basic game loop
// The game loop is as follows:
// 1. Multiple Clients connect to the server

// 2. The client sends a "deal" message to the server. This will need to be triggered by some form of timer instead of a client message
// 2a. The server deals two cards to the client(s)

// 3. The client sends a "com" message to the server. Same thing - next thing to implement is action timers.
// 3a. The server deals three cards to the table

// 4. The client sends a "com" message to the server
// 4a. The server deals one card to the table

// 5. The client sends a "com" message to the server
// 5a. The server deals one card to the table

// Next Ideas:
// - Implement a timer for each action
// - Implement a ready check for each player
// - Store all of this info in the database. Persist the game state.
