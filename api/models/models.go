package models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
)

type Table struct {
	ID             string   `json:"id"`
	Tnum           int      `json:"tnum"`
	MaxPlayers     int      `json:"maxplayers"`
	CurrentPlayers int      `json:"currentplayers"`
	Players        []string `json:"players"`
}

type Gametable struct {
	ID    string   `json:"id"`
	User  string   `json:"user"`
	Table string   `json:"table"`
	Cards []string `json:"cards"`
}

type Game struct {
	ID           string   `json:"id"`
	Table        string   `json:"table"`
	Drawn        []string `json:"drawn"`
	Action       string   `json:"action"`
	Pot          int      `json:"pot"`
	ActionPlayer string   `json:"actionuser"`
}

func PocketBaseClient() *pocketbase.PocketBase {
	pb := pocketbase.New()
	return pb
}

func GetTable(tableid string) (Table, error) {
	pb := PocketBaseClient()
	table := Table{}
	err := pb.Dao().DB().NewQuery("SELECT * FROM table WHERE id = {:id}").Bind(dbx.Params{"id": tableid}).One(&table)
	return table, err
}

// This needs work
