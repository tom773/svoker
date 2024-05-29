package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type Request struct {
	RecordID    string `json:"recordid"`
	TableID     string `json:"tableid"`
	UserID      string `json:"userid"`
	GameID      string `json:"gameid"`
	GametableID string `json:"gametableid"`
}

func SetupRoutes() {
	var app = pocketbase.New()
	var request Request

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		// Pages that hit this:
		e.Router.POST("/api/basicuser", func(c echo.Context) error {

			if err := c.Bind(&request); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}
			response := map[string]interface{}{
				"user": GetUserInfo(*app, request.UserID),
			}

			return c.JSON(http.StatusOK, response)
		})
		// Pages that hit this:
		e.Router.POST("/api/avatar", func(c echo.Context) error {
			if err := c.Bind(&request); err != nil {
				return c.JSON(400, map[string]interface{}{"Error Binding": err.Error()})
			}
			response := map[string]interface{}{
				"avatar": GetAvatar(*app, request.UserID),
			}
			return c.JSON(http.StatusOK, response)
		})
		// Pages that hit this: /holdem/tableid
		e.Router.POST("/api/table/avatar", func(c echo.Context) error {
			if err := c.Bind(&request); err != nil {
				return c.JSON(400, map[string]interface{}{"Error Binding": err.Error()})
			}
			response := getTablePlayers(*app, request.TableID)
			return c.JSON(http.StatusOK, response)
		})
		return nil
	})

	if err := app.Start(); err != nil {
		panic(err)
	}

}

type User struct {
	ID       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Avatar   string `db:"avatar" json:"avatar"`
	Balance  int    `db:"balance" json:"balance"`
}
type CardInfoResponse struct {
	Cards string `json:"cards"`
}
type GameInfoResponse struct {
	GameID      string `db:"id" json:"gameid"`
	GametableID string `db:"id" json:"gametableid"`
	TableID     string `db:"id" json:"tableid"`
	Drawn       string `db:"drawn" json:"drawn"`
}

func GetAvatar(app pocketbase.PocketBase, id string) string {
	user := User{}
	fmt.Println("ID: ", id)
	err := app.Dao().DB().NewQuery("SELECT avatar FROM users WHERE id = {:id}").Bind(dbx.Params{"id": id}).One(&user)
	if err != nil {
		panic(err)
	}
	return user.Avatar
}

// As much as I love websockets, this is just without a doubt the easiest
// way to get the complex data from the database.
func getTablePlayers(app pocketbase.PocketBase, tableid string) []User {

	users := []User{}
	// Finest peice of Sqeual I've written
	err := app.Dao().DB().NewQuery(
		`SELECT users.id, users.username, users.avatar, users.balance 
         FROM users 
         INNER JOIN v2tables 
         ON users.id IN (SELECT value FROM json_each(v2tables.players)) 
         WHERE v2tables.id = {:tableid}`).
		Bind(dbx.Params{"tableid": tableid}).
		All(&users)

	if err != nil {
		panic(err)
	}

	return users
}
func GetUserInfo(app pocketbase.PocketBase, id string) User {
	user := User{}

	err := app.Dao().DB().NewQuery("SELECT id, username, avatar, balance FROM users WHERE id = {:id}").Bind(dbx.Params{"id": id}).One(&user)
	if err != nil {
		panic(err)
	}
	return user
}

func GetUserHand(app pocketbase.PocketBase, id string, tableid string) CardInfoResponse {
	cinfo := CardInfoResponse{}
	err := app.Dao().DB().NewQuery(`SELECT cards FROM gametable WHERE user = {:id} AND "table" = {:tableid}`).Bind(dbx.Params{"id": id, "tableid": tableid}).One(&cinfo)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	return cinfo
}
