package models

type User struct {
	Avatar          string `json:"avatar"`
	Balance         int    `json:"balance"`
	CollectionID    string `json:"collectionId"`
	CollectionName  string `json:"collectionName"`
	Created         string `json:"created"`
	EmailVisibility bool   `json:"emailVisibility"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	Updated         string `json:"updated"`
	Username        string `json:"username"`
	Verified        bool   `json:"verified"`
}

type Game struct {
	ID             string `json:"id"`
	Table_Number   int    `json:"table_number"`
	CollectionName string `json:"collection_name"`
	CreatedAt      string `json:"created"`
	UpdateAt       string `json:"updated"`
}

type GameUser struct {
	ID             string   `json:"id"`
	GameID         string   `json:"game_id"`
	UserID         string   `json:"user_id"`
	CollectionName string   `json:"collection_name"`
	CardsDealt     []string `json:"cards_dealt"`
	CreatedAt      string   `json:"created"`
	UpdateAt       string   `json:"updated"`
}

type ApiResponse struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"perPage"`
	TotalItems int         `json:"totalItems"`
	TotalPages int         `json:"totalPages"`
	Items      interface{} `json:"items"`
}
