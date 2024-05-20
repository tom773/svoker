package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tom773/svoker/api/game"
	weby "github.com/tom773/svoker/api/ws"
)

type Request struct {
	RecordID    string `json:"recordid"`
	TableID     string `json:"tableid"`
	ID          string `json:"userid"`
	GameID      string `json:"gameid"`
	GametableID string `json:"gametableid"`
}

func SetupRoutes() {
	var app = pocketbase.New()
	var request Request

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/ws", func(c echo.Context) error {
			weby.HandleConnection(c.Response().Writer, c.Request(), game.GameState())
			return nil
		})
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		e.Router.POST("/api/basicuser", func(c echo.Context) error {

			if err := c.Bind(&request); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}
			response := map[string]interface{}{
				"user": GetUserInfo(*app, request.ID),
			}

			return c.JSON(http.StatusOK, response)
		})
		e.Router.POST("/api/avatar", func(c echo.Context) error {
			if err := c.Bind(&request); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}
			response := map[string]interface{}{
				"avatar": GetAvatar(*app, request.ID),
			}
			return c.JSON(http.StatusOK, response)
		})
		e.Router.POST("/api/hand", func(c echo.Context) error {
			if err := c.Bind(&request); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}
			response := GetUserHand(*app, request.ID, request.TableID)

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

func GetAvatar(app pocketbase.PocketBase, id string) User {
	user := User{}
	err := app.Dao().DB().NewQuery("SELECT id, avatar FROM users WHERE id = {:id}").Bind(dbx.Params{"id": id}).One(&user)
	if err != nil {
		panic(err)
	}
	return user
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
