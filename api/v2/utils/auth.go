package utils

import (
	"encoding/json"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

func GetKey() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	idenv := os.Getenv("IDENTITY")
	passenv := os.Getenv("PASSWORD")

	payload := map[string]interface{}{
		"identity": idenv,
		"password": passenv,
	}
	adminJSON, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	client := resty.New()
	resp, err := client.R().SetHeader("Content-Type", "application/json").SetBody(adminJSON).Post("http://localhost:8090/api/admins/auth-with-password")
	if err != nil {
		log.Fatal(err)
	}
	var apiResponse map[string]interface{}
	err = json.Unmarshal(resp.Body(), &apiResponse)
	if err != nil {
		log.Fatal(err)
	}
	return apiResponse["token"].(string)
}
