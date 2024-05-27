package client

import (
	"encoding/json"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/tom773/casgo/deck"
	"github.com/tom773/casgo/models"
	"github.com/tom773/casgo/utils"
)

// This sends the full game to the server, which I think works better than having the server handle actions
func Deal(d deck.Deck, t string) {
	client := initClient()
	resp, err := client.R().Get("http://localhost:8090/api/collections/v2gameuser/records?filter(gameID='" + t + "'))")
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
	counter := 0
	for p := range players {
		handPayload := map[string]interface{}{
			"hand": []interface{}{d[counter], d[counter+1]},
		}

		id := players[p]["id"].(string)
		handJSON, err := json.Marshal(handPayload)
		_, err = client.R().SetBody(handJSON).Patch("http://localhost:8090/api/collections/v2gameuser/records/" + id)
		if err != nil {
			log.Fatal(err)
		}
		counter = counter + 2
	}

	tablePayload := map[string]interface{}{
		"deck": d[ctp:],
	}
	tableJSON, err := json.Marshal(tablePayload)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.R().SetBody(tableJSON).Patch("http://localhost:8090/api/collections/v2game/records/" + t)
	if err != nil {
		log.Fatal(err)
	}
}

func get_(coll string, filter string) []map[string]interface{} {
	client := initClient()
	var apiResponse models.ApiResponse

	if filter == "" {
		resp, err := client.R().Get("http://localhost:8090/api/collections/" + coll + "/records")
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(resp.Body(), &apiResponse)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		resp, err := client.R().Get("http://localhost:8090/api/collections/" + coll + "/records?filter(" + filter + ")")
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(resp.Body(), &apiResponse)
		if err != nil {
			log.Fatal(err)
		}
	}

	return returnRaw(apiResponse)
}

func Reset(t string) {
	client := initClient()
	empty := map[string]interface{}{
		"deck": []interface{}{},
	}
	emptyJSON, err := json.Marshal(empty)
	_, err = client.R().SetBody(emptyJSON).Patch("http://localhost:8090/api/collections/v2game/records/" + t)
	if err != nil {
		log.Fatal(err)
	}

	var users []map[string]interface{}
	usersJSON, err := json.Marshal(get_("v2gameuser", "gameID='"+t+"'"))

	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(usersJSON, &users)
	if err != nil {
		log.Fatal(err)
	}
	for user := range users {
		emptyHand := map[string]interface{}{
			"hand": []interface{}{},
		}
		emptyHandJSON, err := json.Marshal(emptyHand)
		_, err = client.R().SetBody(emptyHandJSON).Patch("http://localhost:8090/api/collections/v2gameuser/records/" + users[user]["id"].(string))
		if err != nil {
			log.Fatal(err)
		}
	}
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

func initClient() *resty.Client {
	key := utils.GetKey()
	client := resty.New().
		SetAuthToken(key).
		SetHeader("Content-Type", "application/json")
	return client
}
