package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/tom773/casgo/models"
)

// Onto something here
func V2connect() {
	client := resty.New()
	var apiResponse models.ApiResponse
	resp, err := client.R().
		Get("http://localhost:8090/api/collections/users/records")

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(resp.Body(), &apiResponse)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range apiResponse.Items {
		fmt.Printf("User: %v\n", user.ID)
	}
}
