package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/tom773/casgo/deck"
	"github.com/tom773/casgo/models"
)

// This sends the full game to the server, which I think works better than having the server handle actions
func Deal(d deck.Deck, t string) {
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8090/api/collections/v2gameuser/records?filter(gameID='" + t + "')")
	if err != nil {
		log.Fatal(err)
	}

	var apiResponse models.ApiResponse
	err = json.Unmarshal(resp.Body(), &apiResponse)
	if err != nil {
		log.Fatal(err)
	}

	ctp := (apiResponse.TotalItems) * 2

	players := returnRaw(apiResponse)
	fmt.Println("Full deck: ", d)
	counter := 0
	for p := range players {
		handPayload := map[string]interface{}{
			"hand": []interface{}{d[counter], d[counter+1]},
		}

		id := players[p]["id"].(string)
		handJSON, err := json.Marshal(handPayload)
		_, err = client.R().SetHeader("Content-Type", "application/json").SetBody(handJSON).Patch("http://localhost:8090/api/collections/v2gameuser/records/" + id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, " was dealt ", []interface{}{d[counter], d[counter+1]})
		counter = counter + 2
	}

	tablePayload := map[string]interface{}{
		"deck": d[ctp:],
	}
	fmt.Println("The table was given the rest of the deck: ", d[ctp:])
	tableJSON, err := json.Marshal(tablePayload)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.R().
		SetHeader("Content-Type", "application/json").SetBody(tableJSON).Patch("http://localhost:8090/api/collections/v2game/records/" + t)
	if err != nil {
		log.Fatal(err)
	}
}

func Get(coll string) {
	client := resty.New()

	var apiResponse models.ApiResponse
	resp, err := client.R().Get("http://localhost:8090/api/collections/" + coll + "/records")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(resp.Body(), &apiResponse)
	if err != nil {
		log.Fatal(err)
	}

	Testflex(apiResponse)
}
func returnRaw(response models.ApiResponse) []map[string]interface{} {
	itemsJSON, err := json.Marshal(response.Items)
	if err != nil {
		log.Fatal(err)
	}
	var rawitems []map[string]interface{}
	err = json.Unmarshal(itemsJSON, &rawitems)
	if err != nil {
		log.Fatal(err)
	}
	return rawitems
}
func Testflex(response models.ApiResponse) {
	rawitems := returnRaw(response)

	for _, item := range rawitems {
		coll := item["collectionName"]
		switch coll {
		case "users":
			var user models.User
			mapToStruct(item, &user)
		case "v2game":
			var game models.Game
			mapToStruct(item, &game)
		case "v2gameuser":
			var gameuser models.GameUser
			mapToStruct(item, &gameuser)
		default:
			fmt.Println("Unknown collection")
		}
	}
}

func mapToStruct(data map[string]interface{}, result interface{}) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonData, result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
