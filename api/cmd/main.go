package main

import (
    "encoding/json"
    "net/http"
    "math/rand"
    "time"

    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/models"
    "github.com/pocketbase/pocketbase/plugins/serve"
    "github.com/labstack/echo/v5"
)

type GameState struct {
    TableCards []string     `json:"table_cards"`
    Players    []PlayerState `json:"players"`
}

type PlayerState struct {
    PlayerID string   `json:"player_id"`
    Cards    []string `json:"cards"`
}

type AdvanceGameRequest struct {
    PlayerID string `json:"player_id"`
}

type PlayerResponse struct {
    PlayerCards []string `json:"player_cards"`
    TableCards  []string `json:"table_cards"`
}

func main() {
    app := pocketbase.New()

    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())

    // Register the custom endpoint
    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        e.Router.POST("/api/advance-game", func(c echo.Context) error {
            var request AdvanceGameRequest
            if err := c.Bind(&request); err != nil {
                return c.JSON(http.StatusBadRequest, err.Error())
            }

            // Retrieve the game state from PocketBase
            collection, err := app.Dao().FindCollectionByNameOrId("game_state")
            if err != nil {
                return c.JSON(http.StatusInternalServerError, err.Error())
            }

            gameRecord, err := app.Dao().FindFirstRecordByData(collection, "id", "game_state_id")
            if err != nil {
                return c.JSON(http.StatusInternalServerError, err.Error())
            }

            var gameState GameState
            if err := json.Unmarshal([]byte(gameRecord.GetString("state")), &gameState); err != nil {
                return c.JSON(http.StatusInternalServerError, err.Error())
            }

            // Find the player's state
            var playerState *PlayerState
            for _, player := range gameState.Players {
                if player.PlayerID == request.PlayerID {
                    playerState = &player
                    break
                }
            }

            if playerState == nil {
                return c.JSON(http.StatusNotFound, "Player not found")
            }

            // Example of using a random number in game logic
            randomIndex := rand.Intn(len(gameState.TableCards))
            randomCard := gameState.TableCards[randomIndex]

            // Here you can implement your game logic to advance the state
            // For example, updating the game state with the randomly selected card

            // Return the relevant part of the game state to the player
            response := PlayerResponse{
                PlayerCards: playerState.Cards,
                TableCards:  gameState.TableCards,
            }

            return c.JSON(http.StatusOK, response)
        })
        return nil
    })

    app.Serve(serve.Config{
        PublicDir: "./pb_public",
    })
}
