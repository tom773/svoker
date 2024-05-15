package main

import (
    "os"
    "log"
    "net/http"
    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/apis"
    "github.com/pocketbase/dbx"
    "github.com/labstack/echo/v5"
)

type Request struct {
    RecordID string `json:"recordid"`
    TableID string `json:"tableid"`
    ID string `json:"userid"`
    GameID string `json:"gameid"`
    GametableID string `json:"gametableid"`
}


func main() {
    var app = pocketbase.New()
    var request Request
    // Register the custom endpoint
    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
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
