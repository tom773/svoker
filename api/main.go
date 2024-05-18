package main

import (
    "fmt"
    "os"
    "log"
    "net/http"
    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/apis"
    "github.com/pocketbase/dbx"
    "github.com/labstack/echo/v5"
    "github.com/gorilla/websocket"
    "math/rand/v2"
)
var resetCardDeck = []string{
     "AH", "KH", "QH", "JH", "TH", "9H", "8H", "7H", "6H", "5H", "4H", "3H", "2H",
     "AC", "KC", "QC", "JC", "TC", "9C", "8C", "7C", "6C", "5C", "4C", "3C", "2C",
     "AD", "KD", "QD", "JD", "TD", "9D", "8D", "7D", "6D", "5D", "4D", "3D", "2D",
     "AS", "KS", "QS", "JS", "TS", "9S", "8S", "7S", "6S", "5S", "4S", "3S", "2S",
}

var tableDecks = make(map[string][]string)
    
type Request struct {
    RecordID string `json:"recordid"`
    TableID string `json:"tableid"`
    ID string `json:"userid"`
    GameID string `json:"gameid"`
    GametableID string `json:"gametableid"`
}

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan map[string]interface{})
var gameStateTest = make(map[string]interface{})

func main() {
    var app = pocketbase.New()
    var request Request

    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        e.Router.GET("/ws", func(c echo.Context) error{
            handleConnection(c.Response().Writer, c.Request())
            return nil
        })
        e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
        e.Router.POST("/api/basicuser", func(c echo.Context) error {
            
            if err := c.Bind(&request); err != nil {
                return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
            }
            response := map[string]interface{}{
                "user": getUserInfo(*app, request.ID),
            }

            return c.JSON(http.StatusOK, response)
        })
        e.Router.POST("/api/avatar", func(c echo.Context) error {
            if err := c.Bind(&request); err != nil {
                return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()}) 
            }
            response := map[string]interface{}{
                "avatar": getAvatar(*app, request.ID),
            }
            return c.JSON(http.StatusOK, response)
        })
        e.Router.POST("/api/hand", func(c echo.Context) error { 
            if err := c.Bind(&request); err != nil {
                return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
            }
            response := getUserHand(*app, request.ID, request.TableID)
            
            return c.JSON(http.StatusOK, response)
        })
        return nil
    })
    
    if err := app.Start(); err != nil {
        panic(err)
    }

}

type User struct {
    ID string `db:"id" json:"id"`
    Username string `db:"username" json:"username"`
    Avatar string `db:"avatar" json:"avatar"`
    Balance int `db:"balance" json:"balance"`
}
type CardInfoResponse struct{
    Cards string `json:"cards"`
}
type GameInfoResponse struct {
    GameID string `db:"id" json:"gameid"`
    GametableID string `db:"id" json:"gametableid"`
    TableID string `db:"id" json:"tableid"`
    Drawn string `db:"drawn" json:"drawn"`
}

func getAvatar(app pocketbase.PocketBase, id string) User {
    user := User{}
    err := app.Dao().DB().NewQuery("SELECT id, avatar FROM users WHERE id = {:id}").Bind(dbx.Params{"id": id}).One(&user)
    if err != nil {
        panic(err)
    }
    return user
}
// Balance, Username, and Avatar Information API
func getUserInfo(app pocketbase.PocketBase, id string) User {
    user := User{}

    err := app.Dao().DB().NewQuery("SELECT id, username, avatar, balance FROM users WHERE id = {:id}").Bind(dbx.Params{"id": id}).One(&user)
    if err != nil { 
        panic(err)
    }
    return user
}

func getUserHand(app pocketbase.PocketBase, id string, tableid string) CardInfoResponse {
    cinfo := CardInfoResponse{}
    err := app.Dao().DB().NewQuery(`SELECT cards FROM gametable WHERE user = {:id} AND "table" = {:tableid}`).Bind(dbx.Params{"id": id, "tableid": tableid}).One(&cinfo)    
    if err != nil {
        log.Fatalf("Error %v", err)
    }
    return cinfo 
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    clients[conn] = true
    for {
        var msg map[string]interface{}
        err := conn.ReadJSON(&msg)

        if err != nil {
            log.Fatal(err)
            return
        }
        handleMessage(conn, msg)
    }
}

func handleMessage(ws *websocket.Conn, msg map[string]interface{}) {
    response := make(map[string]interface{})
    tableID, ok := msg["tableID"].(string)
    if !ok {
        response["error"] = "Invalid table ID"
        ws.WriteJSON(response)
        return
    }
    switch msg["type"] {
        case "update":
            name, ok := msg["data"].(string)
            if ok {
                response["message"] = fmt.Sprintf("Hello, %s!", name)
            } else {
                response["error"] = "Invalid name"
            }
        case "deal":
            response["type"] = "dealResponse"
            response["cards"] = drawCard(tableID, 2)
        case "com":
            response["type"] = "comResponse"
            response["cards"] = drawCard(tableID, 5)
        case "reset":
            resetDeck(tableID)
            response["type"] = "resetResponse"
            response["msg"] = "Deck has been reset"
        default:
            response["error"] = "Invalid message type"
    }
    ws.WriteJSON(response)
}

func updateGameState(msg map[string]interface{}) {
    gameStateTest = msg
}

func drawCard(tableID string, n int) []string {
    deck, exists := tableDecks[tableID]
    if !exists {
        deck = resetCardDeck
        tableDecks[tableID] = deck
    }
    cards := make([]string, n)
    for i := 0; i < n; i++ {
        if len(deck) == 0 {
            break
        }
        card := deck[rand.IntN(len(deck))]
        cards[i]= card
        deck = remove(deck, card)
    }
    tableDecks[tableID] = deck
    return cards
}

func remove(s []string, v string) []string {
    i := Index(s, v)
    if i == -1 {
        return s
    }
    return append(s[:i], s[i+1:]...)
}
func Index(s []string, v string) int {
	for i, vs := range s {
		if vs == v {
			return i
		}
	}

	return -1
}

func resetDeck(tableID string) {
    tableDecks[tableID] = resetCardDeck
}
